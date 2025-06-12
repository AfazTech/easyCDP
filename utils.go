package easyCDP

import (
	"os"
	"strings"

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
func (b *Browser) ScreenshotElement(selector string, filename string) error {
	var buf []byte
	err := b.Run(chromedp.Screenshot(selector, &buf, resolveSelector(selector)))
	if err != nil {
		return err
	}
	return os.WriteFile(filename, buf, 0644)
}

func resolveSelector(selector string) chromedp.QueryOption {
	if isXPath(selector) {
		return chromedp.BySearch
	}
	return chromedp.ByQuery
}

func isXPath(selector string) bool {
	return strings.HasPrefix(selector, "/") || strings.HasPrefix(selector, "./")
}

func escapeJSString(str string) string {
	str = strings.ReplaceAll(str, `\`, `\\`)
	str = strings.ReplaceAll(str, `'`, `\'`)
	str = strings.ReplaceAll(str, "\n", `\n`)
	str = strings.ReplaceAll(str, "\r", `\r`)
	return str
}
