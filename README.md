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
- [Features](#features)
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

## Features

* Open and manage tabs
* Navigate pages
* Click elements and send keys
* Take screenshots of full page or specific elements
* Check if elements exist or are visible
* Wait for elements or page load
* Evaluate JavaScript expressions

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