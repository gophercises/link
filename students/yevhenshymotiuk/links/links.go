// Package links provides data structures
// and methods to work with HTML links
package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link provides information about link element
type Link struct {
	Href string
	Text string
}

func text(n *html.Node) string {
	var result strings.Builder

	for c := n.FirstChild; c != nil && c.Data != "a"; c = c.NextSibling {
		var t string

		switch c.Type {
		case html.TextNode:
			t = c.Data
		case html.ElementNode:
			if c.FirstChild != nil {
				t = c.FirstChild.Data
			}
		}

		result.WriteString(t)
	}

	return strings.TrimSpace(result.String())
}

// GetLinks return brief descriptions of every link
// in HTML node
func GetLinks(n *html.Node) []Link {
	var links []Link

	if n.Type == html.ElementNode && n.Data == "a" {
		var href string

		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href = attr.Val
			}
		}

		links = append(links, Link{Href: href, Text: text(n)})
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, GetLinks(c)...)
	}

	return links
}

// GetLinks return brief descriptions of every link
// in HTML file
func GetLinksFromReader(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	return GetLinks(doc), nil
}
