package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	err := ChromedpPrintPdf("http://127.0.0.1:8081/SetDatassetApply.html", "apply.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ChromedpPrintPdf(url string, to string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}

	if err := ioutil.WriteFile(to, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}

	return nil
}
