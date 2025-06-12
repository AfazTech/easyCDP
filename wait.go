package easyCDP

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func (b *Browser) WaitElementTagWithText(tag, text string, timeout time.Duration) (bool, error) {
	textJSON, err := json.Marshal(text)
	if err != nil {
		return false, err
	}
	script := fmt.Sprintf(`
		(function() {
			var elements = document.querySelectorAll('%s');
			for (var i = 0; i < elements.length; i++) {
				if (elements[i].textContent.includes(%s)) {
					return true;
				}
			}
			return false;
		})();
	`, tag, escapeJSString(string(textJSON)))

	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()

	var found bool
	for {
		err := chromedp.Run(ctx, chromedp.Evaluate(script, &found))
		if err != nil {
			return false, err
		}
		if found {
			return true, nil
		}
		select {
		case <-ctx.Done():
			return false, fmt.Errorf("no %s tag containing text '%s' found within timeout", tag, text)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (b *Browser) WaitVisible(selector string, timeout time.Duration) (bool, error) {
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			(function() {
				let el = document.evaluate('%s', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				return el !== null && el.offsetWidth > 0 && el.offsetHeight > 0;
			})();
		`, escapeJSString(selector))
	} else {
		script = fmt.Sprintf(`
			(function() {
				const el = document.querySelector('%s');
				return el !== null && el.offsetWidth > 0 && el.offsetHeight > 0;
			})();
		`, escapeJSString(selector))
	}

	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()

	var visible bool
	for {
		err := chromedp.Run(ctx, chromedp.Evaluate(script, &visible))
		if err != nil {
			return false, err
		}
		if visible {
			return true, nil
		}
		select {
		case <-ctx.Done():
			return false, nil
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (b *Browser) WaitForLoad(timeout time.Duration) (bool, error) {
	ctx, cancel := context.WithTimeout(b.ctx, timeout)
	defer cancel()

	var readyState string
	for {
		err := chromedp.Run(ctx, chromedp.Evaluate(`document.readyState`, &readyState))
		if err != nil {
			return false, err
		}
		if readyState == "complete" {
			return true, nil
		}
		select {
		case <-ctx.Done():
			return false, nil
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func (b *Browser) WaitForJSLoad(timeout time.Duration) (bool, error) {
	time.Sleep(time.Second)
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

func (b *Browser) ElementIsVisible(selector string) (bool, error) {
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			(function() {
				let el = document.evaluate('%s', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				return el !== null && el.offsetWidth > 0 && el.offsetHeight > 0;
			})();
		`, escapeJSString(selector))
	} else {
		script = fmt.Sprintf(`
			(function() {
				const el = document.querySelector('%s');
				return el !== null && el.offsetWidth > 0 && el.offsetHeight > 0;
			})();
		`, escapeJSString(selector))
	}

	var isVisible bool
	err := b.Run(chromedp.Evaluate(script, &isVisible))
	return isVisible, err
}

func (b *Browser) ElementExists(selector string) (bool, error) {
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			(function() {
				return document.evaluate('%s', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue !== null;
			})();
		`, escapeJSString(selector))
	} else {
		script = fmt.Sprintf(`
			(function() {
				return document.querySelector('%s') !== null;
			})();
		`, escapeJSString(selector))
	}

	var exists bool
	err := b.Run(chromedp.Evaluate(script, &exists))
	return exists, err
}

func (b *Browser) WaitExists(selector string, timeout time.Duration) (bool, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		exists, err := b.ElementExists(selector)
		if err != nil {
			return false, err
		}
		if exists {
			return true, nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return false, nil
}

func (b *Browser) WaitNotVisible(selector string, timeout time.Duration) (bool, error) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		exists, err := b.ElementExists(selector)
		if err != nil {
			return false, err
		}
		if !exists {
			return true, nil
		}

		visible, err := b.ElementIsVisible(selector)
		if err != nil {
			return false, err
		}
		if !visible {
			return true, nil
		}

		time.Sleep(100 * time.Millisecond)
	}
	return false, nil
}
