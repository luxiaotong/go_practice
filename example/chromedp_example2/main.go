package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const (
	u = "http://zwy.jtw.beijing.gov.cn/ows/app/common/userevidence/manage?id=b4f09bffb35b47b282fa909d5f852373"
	h = "http://zwy.jtw.beijing.gov.cn"
)

func main() {
	timeCtx, cancel := context.WithTimeout(GetChromeCtx(false), 30*time.Second)
	defer cancel()
	var collectLink string
	err := chromedp.Run(timeCtx,
		chromedp.Navigate(u),
		chromedp.WaitVisible(`//img[@width="850px"]`),
		chromedp.Location(&collectLink),
	)
	if err != nil {
		log.Println("读取失败1: ", err.Error())
		return
	}
	log.Println("正在采集列表: ", collectLink)
	var aLinks []*cdp.Node
	if err := chromedp.Run(timeCtx, chromedp.Nodes(`//img[@width="850px"]`, &aLinks)); err != nil {
		log.Println("读取失败2: ", err.Error())
		return
	}
	for i := 0; i < len(aLinks); i++ {
		src := aLinks[i].AttributeValue("src")
		log.Println("src: ", h+src)
		client := &http.Client{}
		resp, err := client.Get(h + src)
		if err != nil {
			log.Println("读取失败3: ", err.Error())
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("read resp failed: ", err)
			return
		}
		if err := ioutil.WriteFile("img.png", body, 0644); err != nil {
			log.Println("write file failed: ", err)
			return
		}
	}
}

//检查是否有9222端口，来判断是否运行在linux上
func checkChromePort() bool {
	addr := net.JoinHostPort("", "9222")
	conn, err := net.DialTimeout("tcp", addr, 1*time.Second)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

// ChromeCtx 使用一个实例
var ChromeCtx context.Context

func GetChromeCtx(focus bool) context.Context {
	if ChromeCtx == nil || focus {
		allocOpts := chromedp.DefaultExecAllocatorOptions[:]
		allocOpts = append(allocOpts,
			chromedp.DisableGPU,
			chromedp.Flag("blink-settings", "imagesEnabled=false"),
			chromedp.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.55 Safari/537.36`),
			chromedp.Flag("accept-language", `zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7,zh-TW;q=0.6`),
		)
		if checkChromePort() {
			// 不知道为何，不能直接使用 NewExecAllocator ，因此增加 使用 ws://127.0.0.1:9222/ 来调用
			c, _ := chromedp.NewRemoteAllocator(context.Background(), "ws://127.0.0.1:9222/devtools/browser/3483dd92-8aa9-42b4-8544-770c14a20b1a")
			ChromeCtx, _ = chromedp.NewContext(c)
		} else {
			c, _ := chromedp.NewExecAllocator(context.Background(), allocOpts...)
			ChromeCtx, _ = chromedp.NewContext(c)
		}
	}

	return ChromeCtx
}
