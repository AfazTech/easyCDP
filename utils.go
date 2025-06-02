package cdp

import (
	"os"

	"github.com/chromedp/chromedp"
)

func HandleFlags(flags []Flag) []chromedp.ExecAllocatorOption {
	opt := append([]chromedp.ExecAllocatorOption{}, chromedp.DefaultExecAllocatorOptions[:]...)
	for _, flag := range flags {
		opt = append(opt, chromedp.Flag(flag.Key, flag.Value))
	}
	return opt
}

func (b *Browser) Screenshot(filename string) error {
	var buf []byte
	err := b.Run(chromedp.Screenshot("html", &buf, chromedp.NodeVisible, chromedp.ByQuery))
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf, 0644)
}
