// +build ignore

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var rdb *redis.Client
var upgrader = websocket.Upgrader{}
var clients = make(map[*websocket.Conn]bool)

func ws(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	clients[c] = true

	go func() {
		pubsub := rdb.Subscribe(context.Background(), "mychannel1")
		for {
			msg, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				log.Panic("pubsub recv error")
			}
			for c := range clients {
				c.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
			}
		}
	}()

	go func() {
		for {
			mt, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = c.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
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
	http.HandleFunc("/ws", ws)
	http.ListenAndServe(":8080", nil)
}
