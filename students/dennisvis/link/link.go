package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Anchor is a representation of an anchor tag inside a HTML document
type Anchor struct {
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

func nodeToAnchor(n *html.Node) Anchor {
	var href string
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			href = attr.Val
		}
	}
	text := traverseText(n, "")
	return Anchor{href, text}
}

func traverseAnchors(n *html.Node, anchors []Anchor) []Anchor {
	if n.Type == html.ElementNode && n.Data == "a" {
		anchors = append(anchors, nodeToAnchor(n))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		anchors = traverseAnchors(c, anchors)
	}
	return anchors
}

// ParseAnchors takes io.Reader to a HTML document,
// parses the anchor tags in that document,
// and returns those anchors as Anchor instances
func ParseAnchors(r io.Reader) (achors []Anchor, err error) {
	doc, err := html.Parse(r)
	if err != nil {
		return
	}
	achors = traverseAnchors(doc, make([]Anchor, 0))
	return
}
