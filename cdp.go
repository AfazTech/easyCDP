package cdp

import (
	"context"
	"os"

	"github.com/chromedp/chromedp"
)

type Browser struct {
	ctx    context.Context
	cancel context.CancelFunc
}

type Flag struct {
	Key   string
	Value interface{}
}

func NewBrowser(options []Flag) *Browser {
	allocCtx, _ := chromedp.NewExecAllocator(context.Background(), handleFlags(options)...)
	ctx, cancelFunc := chromedp.NewContext(allocCtx)
	return &Browser{
		ctx:    ctx,
		cancel: cancelFunc,
	}
}

func (b *Browser) Navigate(url string) error {
	return chromedp.Run(b.ctx, chromedp.Navigate(url))
}

func handleFlags(flags []Flag) []chromedp.ExecAllocatorOption {
	opt := append(chromedp.DefaultExecAllocatorOptions[:])
	for _, flag := range flags {
		opt = append(opt, chromedp.Flag(flag.Key, flag.Value))
	}
	return opt
}
func (b *Browser) Screenshot(filename string) error {
	var buf []byte
	err := chromedp.Run(b.ctx, chromedp.Screenshot("html", &buf, chromedp.NodeVisible, chromedp.ByQuery))
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf, 0644)
}

func (b *Browser) Close() {
	b.cancel()
}
