package link

import (
	"golang.org/x/net/html"
	"io"
	"strings"
)

// Link is a Href url and its text
type Link struct {
	Href string
	Text string
}

// Parse parses HTML from given io.Reader, returning list of <a href> links found
func Parse(r io.Reader) ([]Link, error) {
	root, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var links []Link
	var rec func(*html.Node)
	rec = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					var text string
					if n.FirstChild != nil {
						text = grabText(n.FirstChild)
					}
					links = append(links, Link{attr.Val, text})
				}
			}
		}
		if n.FirstChild != nil {
			rec(n.FirstChild)
		}
		if n.NextSibling != nil {
			rec(n.NextSibling)
		}
	}
	rec(root)

	return links, nil
}

// grabText returns string of all HTML text starting from given root node, stripping extra whitespace
func grabText(n *html.Node) string {
	var sb strings.Builder
	var rec func(*html.Node)
	rec = func(n *html.Node) {
		if n.Type == html.TextNode {
			s := n.Data
			sb.WriteString(s)
		}
		if n.FirstChild != nil {
			rec(n.FirstChild)
		}
		if n.NextSibling != nil {
			rec(n.NextSibling)
		}
	}
	rec(n)

	return strings.Join(strings.Fields(sb.String()), " ")
}

// LinksString returns reasonable string listing links
func LinksString(links []Link) string {
	var maxW int
	for _, l := range links {
		if len(l.Href) > maxW {
			maxW = len(l.Href)
		}
	}
	maxW++

	var sb strings.Builder
	for _, l := range links {
		sb.WriteString(l.Href)
		for i := 0; i < maxW-len(l.Href); i++ {
			sb.WriteRune(' ')
		}
		sb.WriteString(l.Text)
		sb.WriteRune('\n')
	}

	return sb.String()
}
