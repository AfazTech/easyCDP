# cdp
## cdp is a small project aimed at simplifying browser automation using the chromedp library in Go.

**If this project is helpful to you, you may wish to give it a** :star2: **to support future updates and feature additions!**

### Table of contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Methods](#methods)
- [Todo](#todo)
- [License](#license)
- [Contributors](#contributors)

### Features
- Easy browser initialization with customizable flags
- Check for element existence
- Wait for page load completion
- Check for text visibility on the page
- Take screenshots of the current page
- Manage cookies: get, save, and load cookies
- Equipped with chromedp methods
- Future improvements: Simplifying functions and supporting multiple browser instances

### Prerequisites
- Chrome

### Installation
To install the `cdp` package, run the following command:

```bash
go get github.com/imafaz/cdp
```

### Usage
To use the `cdp` package, you can create a new browser instance and perform various actions. Below is a simple example:

```go
package main

import (
	"log"
	"time"

	"github.com/imafaz/cdp"
)

func main() {
	// Create a new browser 
    
    flags := []cdp.Flag{
		{
			Key:   "headless",
			Value: false,
		},
		{
			Key:   "guest",
			Value: true,
		},
	}
	browser := cdp.NewBrowser(flags)

	// Navigate to a URL
	err := browser.Go("https://example.com")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for the page to load
	_, err = browser.WaitForLoad(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Take a screenshot
	err = browser.Screenshot("screenshot.png")
	if err != nil {
		log.Fatal(err)
	}

	// Close the browser
	browser.Close()
}
```

### Methods
#### `NewBrowser(options []Flag) *Browser`
Creates a new browser instance with the specified options.

#### `Go(url string) error`
Navigates to the specified URL.

#### `Screenshot(filename string) error`
Takes a screenshot of the current page and saves it to the specified filename.

#### `Close()`
Closes the browser instance.

#### `ElementExists(selector string) (bool, error)`
Checks if an element exists on the page.

#### `WaitForLoad(timeout time.Duration) (bool, error)`
Waits for the page to load completely within the specified timeout.

#### `Click(selector string) error`
Clicks on the specified element.

#### `SendKeys(selector string, keys string) error`
Sends keys to the specified input field.

#### `SetValue(selector string, value string) error`
Sets the value of the specified input field.

#### `Evaluate(expression string, res interface{}) error`
Evaluates a JavaScript expression and stores the result in the provided variable.

#### `WaitVisible(selector string, timeout time.Duration) (bool, error)`
Waits for the specified element to become visible within the given timeout duration. Returns true if the element is visible, otherwise returns false.

#### `Reload() error`
Reloads the current page.

#### `GetCookies() ([]*network.Cookie, error)`
Retrieves the cookies from the current browser session.

#### `SaveCookies(filename string) error`
Saves the current cookies to a specified file in JSON format.

#### `LoadCookies(filename string) error`
Loads cookies from a specified file and sets them in the current browser session.

### TODO:
- [ ] Add more comprehensive error handling
- [ ] Implement additional browser actions
- [ ] Improve documentation and examples

### License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

### Contributors
Feel free to contribute to the project by submitting issues or pull requests!