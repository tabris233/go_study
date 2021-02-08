package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		word, image := countWordsAndImages(c)
		words += word
		images += image
	}

	return
}

func main() {
	words, images, err := CountWordsAndImages("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("words: %d images: %d\n", words, images)
}
