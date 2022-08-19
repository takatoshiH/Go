package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

var userAgent string = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36"
var url string = "https://skyticket.jp/"

func main() {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false), // headless=false に変更
		chromedp.Flag("disable-gpu", false),
		chromedp.Flag("enable-automation", true),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("hide-scrollbars", false),
		chromedp.Flag("mute-audio", false),
	)

	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)

	ctx, cancel := chromedp.NewContext(
		allocCtx,
		chromedp.WithLogf(log.Printf),
	)

	//ctx, cancel := chromedp.NewContext(
	//	context.Background(),
	//	chromedp.WithLogf(log.Printf),
	//)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	var value string
	//var imageBuf []byte

	err := chromedp.Run(ctx,
		emulation.SetUserAgentOverride(userAgent),
		chromedp.Navigate(url),
		chromedp.WaitVisible("#footer"),
		chromedp.Text("#js-btnSearchTicket", &value, chromedp.NodeVisible, chromedp.ByQuery),
		chromedp.Click("#js-btnSearchTicket", chromedp.NodeVisible),
		//chromedp.Screenshot("#flights-in-japan", &imageBuf, chromedp.NodeVisible, chromedp.ByQuery),
	)

	if err != nil {
		fmt.Println("例外吐いてるです")
		fmt.Printf("%T", err)
		log.Fatal(err)
	}

	fmt.Println(value)
}
