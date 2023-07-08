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

func (p *Parser) Parse() (*html.Node, error) {
	doc, err := html.Parse(p.Page)
	if err != nil {
		log.Println("error parsing page in Parse()")
		return nil, err
	}
	return doc, nil
}

func (p *Parser) ExtractLinks(doc *html.Node, links *[]Link) {
	traverseDoc(doc, links)
}

func traverseDoc(n *html.Node, links *[]Link) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				text := strings.Trim(n.FirstChild.Data, "\n ")
				if n.FirstChild.NextSibling != nil {
					if n.FirstChild.NextSibling.FirstChild != nil {
						//fmt.Println(n.FirstChild.NextSibling.FirstChild.Data)
						text += " " + strings.Trim(n.FirstChild.NextSibling.FirstChild.Data, "\n ")
					}
				}
				url := a.Val
				*links = append(*links, Link{Url: url, Text: text})
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverseDoc(c, links)
	}
}
