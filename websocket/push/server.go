// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var rdb *redis.Client
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]bool)

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	clients[c] = true

	go func() {
		pubsub := rdb.Subscribe(context.Background(), "screen2")
		for {
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				log.Panic("pubsub recv error")
			}
			log.Printf("channel msg: %v", msg)
			for c := range clients {
				_ = c.SetWriteDeadline(time.Now().Add(3 * time.Second))
				if err = c.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
					log.Print("write message error:", err)
				}
			}
		}
	}()

	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read error:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("echo error:", err)
				break
			}
		}
	}()
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "139.9.119.21:56379",
		Password: "",
		DB:       15,
	})

	fmt.Println("server running...")

	// http.HandleFunc("/ws", ws)
	// http.ListenAndServe(":8080", nil)

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
		return
	}
	s := &http.Server{
		Handler:        &myHandler{},
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-c
		//nolint:gocritic
		// log.Printf("captured %v, stopping profiler and exiting..", sig)
		log.Print("shutdown return")
		_ = lis.Close()
		log.Print("quit app...")
	}()
	cancel := func() {
		c <- os.Interrupt
	}
	if err := s.Serve(lis); err != nil {
		cancel()
		panic(err)
	}
}

type myHandler struct {
}

func (this myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws(w, r)
	log.Print("serve http")
}
