package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// forEachNode针对每个结点x，都会调用pre(x)和post(x)。
// pre和post都是可选的。
// 遍历孩子结点之前，pre被调用
// 遍历孩子结点之后，post被调用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		attr := make([]string, 0, len(n.Attr))
		for _, a := range n.Attr {
			attr = append(attr, a.Key+"=\""+a.Val+"\"")
		}
		tail := ""
		if n.FirstChild == nil {
			tail = " /"
		}
		fmt.Printf("%*s<%s %s%s>\n", depth*2, "", n.Data, strings.Join(attr, " "), tail)
		depth++
	} else if n.Type == html.CommentNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", "[comment]")
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild == nil {
			return
		}
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func getDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return nil, err
	}

	return doc, nil
}

func main() {
	doc, err := getDoc("https://golang.org")

	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	forEachNode(doc, startElement, endElement)
}
