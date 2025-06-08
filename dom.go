package easyCDP

import (
	"fmt"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func (b *Browser) SelectAll(selector string) ([]*cdp.Node, error) {
	var nodes []*cdp.Node
	err := b.Run(chromedp.Nodes(selector, &nodes, resolveSelector(selector)))
	if err != nil {
		return nil, fmt.Errorf("failed to select elements: %v", err)
	}
	return nodes, nil
}

func (b *Browser) Click(selector string) error {
	return b.Run(chromedp.Click(selector, resolveSelector(selector)))
}

func (b *Browser) SendKeys(selector, keys string) error {
	return b.Run(chromedp.SendKeys(selector, keys, resolveSelector(selector)))
}

func (b *Browser) SetValue(selector, value string) error {
	return b.Run(chromedp.SetValue(selector, value, resolveSelector(selector)))
}

func (b *Browser) Evaluate(expression string, res interface{}) error {
	return b.Run(chromedp.Evaluate(expression, res))
}

func (b *Browser) Text(selector string) (string, error) {
	var textContent string
	err := b.Run(chromedp.Text(selector, &textContent, chromedp.NodeVisible, resolveSelector(selector)))
	if err != nil {
		return "", err
	}
	return textContent, nil
}

func (b *Browser) TextExists(text string) (bool, error) {
	var bodyText string
	err := b.Run(chromedp.Text("body", &bodyText, chromedp.NodeVisible, resolveSelector("body")))
	if err != nil {
		return false, err
	}
	return strings.Contains(bodyText, text), nil
}

func (b *Browser) InnerText() (string, error) {
	var bodyText string
	err := b.Run(chromedp.Text("body", &bodyText, chromedp.NodeVisible, resolveSelector("body")))
	if err != nil {
		return "", err
	}
	return bodyText, nil
}

func (b *Browser) ClickTagWithText(tag, text string) error {
	script := fmt.Sprintf(`(function(){
		var elements = document.querySelectorAll('%s');
		for(var i=0; i<elements.length; i++){
			if(elements[i].textContent.includes('%s')){
				elements[i].click();
				return true;
			}
		}
		return false;
	})()`, tag, text)

	var found bool
	err := b.Run(chromedp.Evaluate(script, &found))
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
	err := b.Run(chromedp.OuterHTML("html", &pageSource, resolveSelector("html")))
	if err != nil {
		return "", err
	}
	return pageSource, nil
}

func (b *Browser) GetValue(selector string) (string, error) {
	var value string
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			(function(){
				let result = document.evaluate(%q, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				return result ? result.value : null;
			})()
		`, selector)
	} else {
		script = fmt.Sprintf(`
			(function(){
				let el = document.querySelector(%q);
				return el ? el.value : null;
			})()
		`, selector)
	}
	err := b.Run(chromedp.Evaluate(script, &value))
	if err != nil {
		return "", err
	}
	return value, nil
}

func (b *Browser) Clear(selector string) error {
	return b.Run(chromedp.SetValue(selector, "", resolveSelector(selector)))
}

func (b *Browser) ClickIfExists(selector string) (bool, error) {
	exists, err := b.WaitExists(selector, 2*time.Second)
	if err != nil || !exists {
		return false, err
	}
	err = b.Run(chromedp.Click(selector, resolveSelector(selector)))
	if err != nil {
		return false, err
	}
	return true, nil
}

func (b *Browser) GetAttribute(selector, attr string) (string, error) {
	var value string
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			(function(){
				let node = document.evaluate(%q, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				return node ? node.getAttribute(%q) : null;
			})()
		`, selector, attr)
	} else {
		script = fmt.Sprintf(`
			(function(){
				let el = document.querySelector(%q);
				return el ? el.getAttribute(%q) : null;
			})()
		`, selector, attr)
	}
	err := b.Run(chromedp.Evaluate(script, &value))
	if err != nil {
		return "", err
	}
	return value, nil
}

func (b *Browser) ScrollTo(selector string) error {
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			(function(){
				let node = document.evaluate(%q, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
				if(node) node.scrollIntoView({behavior: "smooth", block: "center"});
			})()
		`, selector)
	} else {
		script = fmt.Sprintf(`
			(function(){
				let el = document.querySelector(%q);
				if(el) el.scrollIntoView({behavior: "smooth", block: "center"});
			})()
		`, selector)
	}
	return b.Run(chromedp.Evaluate(script, nil))
}

func (b *Browser) WaitAndClick(selector string, timeout time.Duration) error {
	exists, err := b.WaitExists(selector, timeout)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("element %s not found within timeout", selector)
	}
	return b.Run(chromedp.Click(selector, resolveSelector(selector)))
}

func (b *Browser) Focus(selector string) error {
	return b.Run(chromedp.Focus(selector, resolveSelector(selector)))
}

func (b *Browser) SetInnerHTML(selector string, html string) error {
	var script string
	if isXPath(selector) {
		script = fmt.Sprintf(`
			let node = document.evaluate(%q, document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue;
			if(node) node.innerHTML = %q;
		`, selector, html)
	} else {
		script = fmt.Sprintf(`
			let el = document.querySelector(%q);
			if(el) el.innerHTML = %q;
		`, selector, html)
	}
	return b.Run(chromedp.Evaluate(script, nil))
}

func (b *Browser) ScrollIntoView(selector string) error {
	return b.Run(chromedp.ScrollIntoView(selector, resolveSelector(selector)))
}
