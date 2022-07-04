package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"github.com/go-ping/ping"
)

const (
	hostURL  = "114.114.114.114"
	pingSize = 4
)

func main() {
	// pinger, err := ping.NewPinger("114.114.114.114")
	// if err != nil {
	// 	panic(err)
	// }
	// pinger.Count = 3
	// err = pinger.Run() // Blocks until finished.
	// if err != nil {
	// 	panic(err)
	// }
	// stats := pinger.Statistics()
	// log.Printf("stats: %#v\n", stats)
	for {
		if err := checkNetworkImplement(hostURL, pingSize); err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}
}

// checkNetworkImplement 执行Ping，并不间断
//
//  param: string host ping的address
//  param: int    size 每几次结果计算网络是否通
func checkNetworkImplement(host string, size int) error {
	sendN := 0
	pings := make([]bool, size)
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return err
	}
	sysType := runtime.GOOS
	if sysType == "windows" {
		pinger.SetPrivileged(true)
	}
	lock := sync.Mutex{}
	pinger.OnRecv = func(pkt *ping.Packet) {
		lock.Lock()
		defer lock.Unlock()
		log.Printf("receive %d %d\n", sendN, pkt.Seq)
		// updatePings(pings, size, sendN, pkt.Seq, true)
		sendN = pkt.Seq
	}
	pinger.Interval = 5 * time.Second

	log.Printf("PING %s (%s)\n", pinger.Addr(), pinger.IPAddr())
	go func() {
		time.Sleep(200 * time.Millisecond)
		interval := time.NewTicker(pinger.Interval)
		defer interval.Stop()
		for range interval.C {
			lock.Lock()
			if sendN < pinger.PacketsSent-2 {
				// updatePings(pings, size, sendN, pinger.PacketsSent-1, false)
				sendN = pinger.PacketsSent - 1
				log.Printf("send: %v %v", sendN, pings)
			}
			lock.Unlock()
		}
	}()
	if err := pinger.Run(); err != nil {
		log.Println("pinger run error: ", err)
	}
	return nil
}
