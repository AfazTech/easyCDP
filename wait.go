package cdp

import (
	"fmt"
	"time"

	"github.com/chromedp/chromedp"
)

func (b *Browser) WaitElementTagWithText(tag, text string, timeout time.Duration) (bool, error) {
	script := fmt.Sprintf(`
		(function() {
			var elements = document.querySelectorAll('%s');
			for (var i = 0; i < elements.length; i++) {
				if (elements[i].textContent.includes('%s')) {
					return true;
				}
			}
			return false;
		})();
	`, tag, text)

	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		var found bool
		err := b.Run(
			chromedp.Evaluate(script, &found),
		)
		if err != nil {
			return false, err
		}

		if found {
			return true, nil
		}

		time.Sleep(100 * time.Millisecond)
	}

	return false, fmt.Errorf("no %s tag containing text '%s' found within timeout", tag, text)
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

func (b *Browser) ElementIsVisible(selector string) (bool, error) {
	var isDisplayed bool
	err := b.Run(chromedp.Evaluate(fmt.Sprintf(`
	(function() {
		const element = document.querySelector('%s');
		if (!element) {
			return false; 
		}
		return element.offsetWidth > 0 && element.offsetHeight > 0;
	})();
	`, selector), &isDisplayed))
	if err != nil {
		return false, err
	}
	return isDisplayed, nil
}

func (b *Browser) ElementExists(selector string) (bool, error) {
	var exists bool
	err := b.Run(chromedp.Evaluate(fmt.Sprintf(`
	(function() {
    const element = document.querySelector('%s');
    return element !== null;
})();
	`, selector), &exists))
	if err != nil {
		return false, err
	}
	return exists, nil
}
