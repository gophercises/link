package links

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

var ReLinks []Link

func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	search(doc)
	return ReLinks, nil
}

func search(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		l := Link{
			Href: n.Attr[0].Val,
			Text: n.FirstChild.Data,
		}
		ReLinks = append(ReLinks, l)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		search(c)
	}
}
