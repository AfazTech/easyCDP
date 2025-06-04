package easyCDP

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func (b *Browser) CaptureNetworkRequests(timeout time.Duration) ([]*network.EventRequestWillBeSent, error) {
	requests := make(chan *network.EventRequestWillBeSent, 100)
	chromedp.ListenTarget(b.ctx, func(ev interface{}) {
		if ev, ok := ev.(*network.EventRequestWillBeSent); ok {
			requests <- ev
		}
	})
	err := b.Run(chromedp.ActionFunc(func(ctx context.Context) error {
		return network.Enable().Do(ctx)
	}))
	if err != nil {
		return nil, fmt.Errorf("failed to enable network events: %w", err)
	}

	time.Sleep(timeout)
	close(requests)

	var capturedRequests []*network.EventRequestWillBeSent
	for req := range requests {
		capturedRequests = append(capturedRequests, req)
	}

	return capturedRequests, nil
}

func (b *Browser) CaptureNetworkRequestsStream() (chan *network.EventRequestWillBeSent, chan error, error) {
	reqChan := make(chan *network.EventRequestWillBeSent, 100)
	errChan := make(chan error, 10)

	err := b.Run(network.Enable())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to enable network events: %w", err)
	}

	chromedp.ListenTarget(b.ctx, func(ev interface{}) {
		switch e := ev.(type) {
		case *network.EventRequestWillBeSent:
			select {
			case reqChan <- e:
			default:
				errChan <- fmt.Errorf("request channel full, dropping request: %s", e.Request.URL)
			}
		case *network.EventLoadingFailed:
			select {
			case errChan <- fmt.Errorf("request failed: %s (%s)", e.ErrorText, e.RequestID.String()):
			default:
			}
		}
	})

	return reqChan, errChan, nil
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
