package cdp

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"

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

func (b *Browser) Go(url string) error {
	return b.Run(chromedp.Navigate(url))
}

func handleFlags(flags []Flag) []chromedp.ExecAllocatorOption {
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

func (b *Browser) Close() {
	b.cancel()
}
func (b *Browser) ElementExists(selector string) (bool, error) {
	var nodes []*cdp.Node
	err := b.Run(chromedp.Nodes(selector, &nodes, chromedp.NodeVisible, chromedp.ByQuery))
	if err != nil {
		return false, err
	}
	return len(nodes) > 0, nil
}

func (b *Browser) Run(actions ...chromedp.Action) error {
	return chromedp.Run(b.ctx, actions...)
}
func (b *Browser) WaitForLoad(timeout time.Duration) (bool, error) {
	time.Sleep(time.Second * 1)
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		var readyState string
		err := b.Run(chromedp.Evaluate("document.readyState", &readyState))
		if err != nil {
			return false, err
		}

		if readyState == "complete" {
			return true, nil
		}

		time.Sleep(1 * time.Second)
	}

	return false, nil
}

func (b *Browser) WaitForJSLoad(timeout time.Duration) (bool, error) {
	time.Sleep(time.Second * 1)
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		var onloadExecuted bool
		err := b.Run(chromedp.Evaluate("window.onload !== null", &onloadExecuted))
		if err != nil {
			return false, err
		}

		if onloadExecuted {
			return true, nil
		}

		time.Sleep(100 * time.Millisecond)
	}

	return false, nil
}

func (b *Browser) Click(selector string) error {
	return b.Run(chromedp.Click(selector, chromedp.ByQuery))
}
func (b *Browser) SendKeys(selector string, keys string) error {

	return b.Run(chromedp.SendKeys(selector, keys, chromedp.ByQuery))
}
func (b *Browser) SetValue(selector string, value string) error {
	return b.Run(chromedp.SetValue(selector, value, chromedp.ByQuery))
}
func (b *Browser) Evaluate(expression string, res interface{}) error {
	return b.Run(chromedp.Evaluate(expression, res))
}
func (b *Browser) Reload() error {
	return b.Run(chromedp.Reload())
}
func (b *Browser) WaitVisible(selector string, timeout time.Duration) (bool, error) {
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		var exists bool
		err := b.Run(chromedp.Evaluate(fmt.Sprintf("document.querySelector('%s') !== null", selector), &exists))
		if err != nil {
			return false, err
		}

		if exists {
			var isVisible bool
			err = b.Run(chromedp.Evaluate(fmt.Sprintf("document.querySelector('%s').offsetWidth > 0 && document.querySelector('%s').offsetHeight > 0", selector, selector), &isVisible))
			if err != nil {
				return false, err
			}

			if isVisible {
				return true, nil
			}
		}

		time.Sleep(100 * time.Millisecond)
	}

	return false, nil
}

func (b *Browser) GetCookies() ([]*network.Cookie, error) {
	var cookies []*network.Cookie
	err := b.Run(chromedp.ActionFunc(func(ctx context.Context) error {
		var err error
		cookies, err = network.GetCookies().Do(ctx)
		return err
	}))
	if err != nil {
		return nil, err
	}
	return cookies, nil
}

func (b *Browser) SaveCookies(filename string) error {
	cookies, err := b.GetCookies()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (b *Browser) LoadCookies(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	var cookies []*network.Cookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		return err
	}

	cookieParams := make([]*network.CookieParam, len(cookies))
	for i, cookie := range cookies {
		var expires *cdp.TimeSinceEpoch
		if cookie.Expires > 0 {
			exp := cdp.TimeSinceEpoch(time.Unix(int64(cookie.Expires), 0))
			expires = &exp
		}

		cookieParams[i] = &network.CookieParam{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Domain:   cookie.Domain,
			Path:     cookie.Path,
			Expires:  expires,
			Secure:   cookie.Secure,
			HTTPOnly: cookie.HTTPOnly,
			SameSite: cookie.SameSite,
		}
	}

	return b.Run(network.SetCookies(cookieParams))
}
func (b *Browser) Text(selector string) (string, error) {
	var textContent string
	err := b.Run(chromedp.Text(selector, &textContent, chromedp.NodeVisible, chromedp.ByQuery))
	if err != nil {
		return "", err
	}
	return textContent, nil
}
func (b *Browser) TextExists(text string) (bool, error) {
	var bodyText string
	err := b.Run(chromedp.Text("body", &bodyText, chromedp.NodeVisible, chromedp.ByQuery))
	if err != nil {
		return false, err
	}
	return strings.Contains(bodyText, text), nil
}

func (b *Browser) InnerText() (string, error) {
	var bodyText string
	err := b.Run(chromedp.Text("body", &bodyText, chromedp.NodeVisible, chromedp.ByQuery))
	if err != nil {
		return "", err
	}
	return bodyText, nil
}
func (b *Browser) GetContext() context.Context {
	return b.ctx
}
func (b *Browser) ClickTagWithText(tag, text string) error {
	script := fmt.Sprintf(`
		(function() {
			var elements = document.querySelectorAll('%s');
			for (var i = 0; i < elements.length; i++) {
				if (elements[i].textContent.includes('%s')) {
					elements[i].click();
					return true;
				}
			}
			return false;
		})();
	`, tag, text)
	var found bool
	err := b.Run(
		chromedp.Evaluate(script, &found),
	)
	if err != nil {
		return err
	}

	if !found {
		return errors.New(fmt.Sprintf("no %s tag containing text '%s' found", tag, text))
	}

	return nil
}
