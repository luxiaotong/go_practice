package main

import (
	"log"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/go-ping/ping"
)

var (
	pingStatus atomic.Value
	pingOnce   sync.Once
	pingLock   sync.Mutex
	sendN      uint64
	recvN      uint64
)

const (
	hostURL  = "114.114.114.114"
	pingSize = 4
	pingIntv = 5 * time.Second
)

func StartPing() {
	pingOnce.Do(func() {
		go func() {
			for {
				if err := checkNetworkImplement(hostURL, pingSize); err == nil {
					break
				}
				time.Sleep(5 * time.Second)
			}
		}()
	})
}

func Status() bool {
	if pingStatus.Load() == nil {
		return false
	}
	return pingStatus.Load().(bool)
}

// checkNetworkImplement 执行Ping，并不间断
//
//  param: string host ping的address
//  param: int    size 每几次结果计算网络是否通
func checkNetworkImplement(host string, size int) error {
	sysType := runtime.GOOS
	pings := make([]bool, size)
	go func() {
		time.Sleep(200 * time.Millisecond)
		interval := time.NewTicker(pingIntv)
		defer interval.Stop()
		for range interval.C {
			pingLock.Lock()
			// log.Debug("ping send: %v recv: %v", sendN, recvN)
			if recvN+2 < sendN {
				updatePings(pings, size, recvN, sendN, false)
				//log.Debug("send: %v %v", sendN, pings)
			}
			pingLock.Unlock()
		}
	}()
	go func() {
		onRecv := func(pkt *ping.Packet) {
			pingLock.Lock()
			defer pingLock.Unlock()
			updatePings(pings, size, recvN, sendN, true)
			atomic.StoreUint64(&recvN, sendN)
			log.Printf("count on recv: %v, send: %v", recvN, sendN)
		}
		interval := time.NewTicker(pingIntv)
		defer interval.Stop()
		for range interval.C {
			pinger, err := ping.NewPinger(host)
			if err != nil {
				log.Printf("new pinger error: %v", err)
				continue
			}
			//log.Debug("PING %s (%s)", pinger.Addr(), pinger.IPAddr())
			if sysType == "windows" {
				pinger.SetPrivileged(true)
			}
			pinger.OnRecv = onRecv
			pinger.Count = 1
			pinger.Timeout = 500 * time.Millisecond
			atomic.AddUint64(&sendN, 1)
			log.Printf("on send seq: %d", sendN)
			if err := pinger.Run(); err != nil {
				log.Printf("ping run error: %v", err)
			}
			log.Printf("on send seq2: %d", sendN)
		}
	}()
	return nil
}

func updatePings(pings []bool, size int, prev, current uint64, result bool) {
	// log.Debug("update pings: %v, size: %v, prev: %v, curr: %v, res: %v", pings, size, prev, current, result)
	for i := prev; i < current; i++ {
		pings[int(i)%size] = result
	}
	// 只要size内有一个true，那么就返回true
	for _, st := range pings {
		if st {
			pingStatus.Store(st)
			return
		}
	}
	pingStatus.Store(false)
}

func main() {
	StartPing()

	time.Sleep(200 * time.Millisecond)
	interval := time.NewTicker(time.Second)
	defer interval.Stop()
	for range interval.C {
		log.Println("internet: ", Status())
	}
}
