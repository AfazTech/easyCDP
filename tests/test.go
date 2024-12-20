package main

import "github.com/imafaz/cdp"

func main() {
	br := cdp.NewBrowser(false)
	br.Navigate("https://hetzner.com")
}
