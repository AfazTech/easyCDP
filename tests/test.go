package main

import (
	"sync"

	"github.com/imafaz/cdp"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	flags := []cdp.Flag{
		{
			Key:   "headless",
			Value: false,
		},
		{
			Key:   "guest",
			Value: true,
		},
		{
			Key:   "window-size",
			Value: "1920,1080",
		},
	}

	// go func() {
	// 	defer wg.Done()
	// 	br := cdp.NewBrowser(flags)

	// 	br.Navigate("https://hetzner.com")

	// }()
	go func() {
		defer wg.Done()
		br := cdp.NewBrowser(flags)

		go br.Navigate("https://ovh.com")

	}()
	wg.Wait()
}
