package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var htmlFile = flag.String("htmlFile", "", "Thw HTML file to parse for links")

type anchor struct {
	Href string
	Text string
}

func traverseText(n *html.Node, text string) string {
	if n.Type == html.TextNode {
		if len(text) > 0 && !strings.HasSuffix(text, " ") {
			text = text + " "
		}
		text = text + strings.TrimSpace(strings.Trim(n.Data, "\n"))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = traverseText(c, text)
	}
	return strings.TrimSpace(text)
}

func nodeToAnchor(n *html.Node) anchor {
	var href string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			href = attr.Val
		}
	}
	text := traverseText(n, "")
	return anchor{href, text}
}

func traverseAnchors(n *html.Node, anchors []anchor) []anchor {
	if n.Type == html.ElementNode && n.Data == "a" {
		anchors = append(anchors, nodeToAnchor(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		anchors = traverseAnchors(c, anchors)
	}
	return anchors
}

func main() {
	flag.Parse()
	if *htmlFile == "" {
		panic("A valid HTML file name is required")
	}

	f, err := os.Open(*htmlFile)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	anchors := traverseAnchors(doc, make([]anchor, 0))

	formatLine := "----------------------------------------"
	fmt.Println(formatLine)
	for _, anchor := range anchors {
		fmt.Printf("href: %s\ntext: %s\n", anchor.Href, anchor.Text)
		fmt.Println(formatLine)
	}
}
