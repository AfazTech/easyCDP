package easyCDP

import (
	"fmt"

	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"
)

func (b *Browser) Reload() error {
	return b.Run(chromedp.Reload())
}

func (b *Browser) GetUrl() (string, error) {
	var url string
	err := b.Run(
		chromedp.Location(&url),
	)
	if err != nil {
		return "", fmt.Errorf("failed to get URL: %v", err)
	}
	return url, nil
}

func (b *Browser) GetTabs() ([]*target.Info, error) {
	return chromedp.Targets(b.ctx)
}

func (b *Browser) Navigate(url string) error {
	return b.Run(chromedp.Navigate(url))
}
func (b *Browser) SwitchTab(tabID target.ID) error {
	Newctx, cancelTab := chromedp.NewContext(b.GetContext(), chromedp.WithTargetID(tabID))
	newb := AttachContext(Newctx, cancelTab)
	err := chromedp.Run(newb.GetContext(), target.ActivateTarget(tabID))
	if err != nil {
		return fmt.Errorf("failed switch tab: %v", err)
	}
	b.SetContext(newb.ctx, newb.cancel)
	return nil
}

func (b *Browser) GetTab(tabID target.ID) (*Browser, error) {
	newCtx, cancelTab := chromedp.NewContext(b.GetContext(), chromedp.WithTargetID(tabID))
	newb := AttachContext(newCtx, cancelTab)
	err := chromedp.Run(newb.GetContext(), target.ActivateTarget(tabID))
	if err != nil {
		return nil, fmt.Errorf("failed get tab: %v", err)
	}
	return newb, nil
}
func (b *Browser) NewTab() (*Browser, error) {
	newCtx, cancelTab := chromedp.NewContext(b.GetContext())

	newBrowser := AttachContext(newCtx, cancelTab)
	err := newBrowser.Navigate("about:blank")
	if err != nil {
		cancelTab()
		return nil, fmt.Errorf("failed to create new tab: %v", err)
	}

	targets, err := chromedp.Targets(newCtx)
	if err != nil {
		cancelTab()
		return nil, fmt.Errorf("failed to get targets: %v", err)
	}

	var newTabID target.ID
	for _, t := range targets {
		if t.Type == "page" && t.URL == "about:blank" {
			newTabID = t.TargetID
			break
		}
	}

	if newTabID == "" {
		cancelTab()
		return nil, fmt.Errorf("new tab not found")
	}

	err = chromedp.Run(newCtx, target.ActivateTarget(newTabID))
	if err != nil {
		cancelTab()
		return nil, fmt.Errorf("failed to activate new tab: %v", err)
	}
	return newBrowser, nil
}
