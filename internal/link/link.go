package link

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
)

type Parser struct {
	Page *os.File
}

func NewParser(pagePath string) (*Parser, error) {
	f, err := os.OpenFile(pagePath, os.O_RDWR|os.O_CREATE, 0755)
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

func (p *Parser) PrintDoc(doc *html.Node) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		//fmt.Println(n.Type, n.Data, n.NextSibling)
		//if n.PrevSibling != nil && n.PrevSibling.Data == "a" {
		//	fmt.Println(n.Data)
		//}
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				//fmt.Println(a.Val, a.Namespace, a.Key)
				if a.Key == "href" {
					fmt.Println(strings.Trim(n.FirstChild.Data, "\n "))
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
