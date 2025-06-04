# easyCDP

## Overview

easyCDP is a lightweight and simple Go library for controlling headless Chrome using the Chrome DevTools Protocol (CDP) via chromedp.  
It provides a clean and easy-to-use API for browser automation tasks.

## Donate

<a href="http://www.coffeete.ir/afaz">
  <img src="http://www.coffeete.ir/images/buttons/lemonchiffon.png" width="260" />
</a>

## Table of Contents

- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)
- [Methods](#methods)
- [TODO](#todo)
- [Contributing](#contributing)
- [License](#license)

## Requirements

- Go 1.21 or higher  
- Chrome or Chromium installed  
- Any OS (tested on Arch Linux (I use arch btw))

## Installation

```bash
go get github.com/AfazTech/easyCDP
````

## Usage

Here is a simple example of how to use easyCDP:

```go
package main

import (
	"log"
	"time"

	"github.com/AfazTech/easyCDP"
)

func main() {
	flags := []easyCDP.Flag{
		{Key: "headless", Value: false},
	}
	browser := easyCDP.NewBrowser(flags)

	defer browser.CloseBrowser()

	tab1, err := browser.NewTab()
	err = tab1.Navigate("https://news.ycombinator.com")
	if err != nil {
		log.Fatal(err)
	}

	visible, err := tab1.WaitVisible("a", 5*time.Second)
	if err != nil || !visible {
		log.Fatal("element not visible")
	}

	text, err := tab1.Text(`a[href="news"]`)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Tag Text:", text)
}
```

## Methods

* `AttachContext(ctx context.Context, cancel context.CancelFunc) *Browser`
* `NewBrowser(options []Flag) *Browser`
* `NewRemoteBrowser(debuggingURL string) (*Browser, error)`
* `(*Browser) CaptureNetworkRequests(timeout time.Duration) ([]*network.EventRequestWillBeSent, error)`
* `(*Browser) CaptureNetworkRequestsStream() (chan *network.EventRequestWillBeSent, chan error, error)`
* `(*Browser) Clear(selector string) error`
* `(*Browser) Click(selector string) error`
* `(*Browser) ClickIfExists(selector string) (bool, error)`
* `(*Browser) ClickTagWithText(tag, text string) error`
* `(*Browser) CloseBrowser()`
* `(*Browser) CloseTab() error`
* `(*Browser) ElementExists(selector string) (bool, error)`
* `(*Browser) ElementIsVisible(selector string) (bool, error)`
* `(*Browser) Evaluate(expression string, res interface{}) error`
* `(*Browser) Focus(selector string) error`
* `(*Browser) GetAttribute(selector, attr string) (string, error)`
* `(*Browser) GetContext() context.Context`
* `(*Browser) GetCookies() ([]*network.Cookie, error)`
* `(*Browser) GetPageSource() (string, error)`
* `(*Browser) GetTab(tabID target.ID) (*Browser, error)`
* `(*Browser) GetTabs() ([]*target.Info, error)`
* `(*Browser) GetUrl() (string, error)`
* `(*Browser) GetValue(selector string) (string, error)`
* `(*Browser) InnerText() (string, error)`
* `(*Browser) LoadCookies(filename string) error`
* `(*Browser) Navigate(url string) error`
* `(*Browser) NewTab() (*Browser, error)`
* `(*Browser) Reload() error`
* `(*Browser) Run(actions ...chromedp.Action) error`
* `(*Browser) SaveCookies(filename string) error`
* `(*Browser) Screenshot(filename string) error`
* `(*Browser) ScreenshotElement(selector string, filename string) error`
* `(*Browser) ScrollIntoView(selector string) error`
* `(*Browser) ScrollTo(selector string) error`
* `(*Browser) SendKeys(selector, keys string) error`
* `(*Browser) SetContext(ctx context.Context, cancel context.CancelFunc)`
* `(*Browser) SetInnerHTML(selector string, html string) error`
* `(*Browser) SetValue(selector, value string) error`
* `(*Browser) SwitchTab(tabID target.ID) error`
* `(*Browser) Text(selector string) (string, error)`
* `(*Browser) TextExists(text string) (bool, error)`
* `(*Browser) WaitAndClick(selector string, timeout time.Duration) error`
* `(*Browser) WaitElementTagWithText(tag, text string, timeout time.Duration) (bool, error)`
* `(*Browser) WaitExists(selector string, timeout time.Duration) (bool, error)`
* `(*Browser) WaitForJSLoad(timeout time.Duration) (bool, error)`
* `(*Browser) WaitForLoad(timeout time.Duration) (bool, error)`
* `(*Browser) WaitNotVisible(selector string, timeout time.Duration) (bool, error)`
* `(*Browser) WaitVisible(selector string, timeout time.Duration) (bool, error)`

## TODO

* Add full network interception support
* Improve documentation and add more examples
* Add more browser control features

## Contributing

Contributions are welcome!

1. Fork the repository
2. Commit your changes (e.g. `git commit -m "Add feature"`)
3. Push your branch (`git push origin feature-branch`)
4. Open a Pull Request

Please add tests or update examples if applicable.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/AfazTech/easyCDP/blob/main/LICENSE) file for details.
