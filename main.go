package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/fenriz07/link/helpers"
	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func main() {

	html := getHtml()

	parseHtml(html)

}

func parseHtml(htmlstring string) {

	doc, err := html.Parse(strings.NewReader(htmlstring))

	if err != nil {
		log.Fatal(err)
	}

	links := []Link{}

	anchor := analyzeNode(doc, &links)

	helpers.DD(anchor)

}

func getHref(attr []html.Attribute) (string, error) {

	//helpers.DD(attr[0].Key)

	for _, a := range attr {

		if a.Key == "href" {
			return a.Val, nil
		}
	}

	return "", errors.New("No se encontro el href")
}

func analyzeNode(n *html.Node, links *[]Link) *[]Link {

	if n.Type == html.ElementNode && n.Data == "a" {

		href, err := getHref(n.Attr)

		if err != nil {
			helpers.DD(err)
		}

		text := ""

		textNode := n.FirstChild

		if textNode != nil {
			text = textNode.Data
		}

		*links = append(*links, Link{href, text})
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		analyzeNode(c, links)
	}

	return links

}

func getHtml() string {
	namefile := flag.String("namefile", "ex1.html", "define to html file to parse")

	flag.Parse()

	bshtml, err := ioutil.ReadFile(*namefile)

	if err != nil {
		helpers.Exit(err)
	}

	return string(bshtml)

}
