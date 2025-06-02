package cdp

import (
	"fmt"
	"strings"

	"github.com/chromedp/chromedp"
)

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
		return fmt.Errorf("no %s tag containing text '%s' found", tag, text)
	}

	return nil
}

func (b *Browser) GetPageSource() (string, error) {
	var pageSource string
	err := b.Run(chromedp.OuterHTML("html", &pageSource, chromedp.ByQuery))
	if err != nil {
		return "", err
	}
	return pageSource, nil
}

func (b *Browser) GetValue(selector string) (string, error) {
	var value string
	script := fmt.Sprintf("document.querySelector('%s').value", selector)
	err := b.Run(chromedp.Evaluate(script, &value))
	if err != nil {
		return "", err
	}
	return value, nil
}
