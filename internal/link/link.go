package link

import (
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Link struct {
	Url  string
	Text string
}

type Parser struct {
	Page *os.File
}

func NewParser(pagePath string) (*Parser, error) {
	f, err := os.OpenFile(pagePath, os.O_RDONLY, 0755)
	if err != nil {
		log.Println("error reading html file")
		return nil, err
	}
	return &Parser{Page: f}, nil
}

func (p *Parser) parse() (*html.Node, error) {
	doc, err := html.Parse(p.Page)
	if err != nil {
		log.Println("error parsing page in Parse()")
		return nil, err
	}
	return doc, nil
}

func (p *Parser) ExtractLinks() ([]Link, error) {
	doc, err := p.parse()
	if err != nil {
		return nil, err
	}
	// 1. Find <a> nodes in doc
	// 2. for each link node...
	//	2a. build a link
	// 3. return the links
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}
	return links, nil
}

func buildLink(n *html.Node) Link {
	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Url = attr.Val
			break
		}
	}
	ret.Text = text(n)
	return ret
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c) + " "
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNodes(c)...)
	}
	return ret
}
