package easyCDP

import (
	"context"

	"github.com/chromedp/cdproto/page"
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

func AttachContext(ctx context.Context, cancel context.CancelFunc) *Browser {
	return &Browser{
		ctx:    ctx,
		cancel: cancel,
	}
}

func NewRemoteBrowser(debuggingURL string) (*Browser, error) {
	ctx, cancelFunc := chromedp.NewRemoteAllocator(context.Background(), debuggingURL)

	return &Browser{
		ctx:    ctx,
		cancel: cancelFunc,
	}, nil
}

func NewBrowser(options []Flag) *Browser {
	ctx, cancelFunc := chromedp.NewExecAllocator(context.Background(), HandleFlags(options)...)
	return &Browser{
		ctx:    ctx,
		cancel: cancelFunc,
	}
}

func (b *Browser) GetContext() context.Context {
	return b.ctx
}

func (b *Browser) CloseTab() error {

	return chromedp.Run(b.GetContext(), page.Close())
}
func (b *Browser) CloseBrowser() {
	if b.cancel != nil {
		b.cancel()
		b.cancel = nil
	}
}

func (b *Browser) Run(actions ...chromedp.Action) error {
	return chromedp.Run(b.ctx, actions...)
}

func (b *Browser) SetContext(ctx context.Context, cancel context.CancelFunc) {
	if b.cancel != nil {
		b.cancel()
	}
	b.ctx = ctx
	b.cancel = cancel
}
