package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"golang.org/x/net/html"
)

var verbose bool

type Link struct {
	Href, Text string
}

// Recursive : parse the document (n should be the root node to parse all)
func parse(n *html.Node, links *[]Link) {
	switch n.Type {
	case html.ElementNode:
		displayNode("Element", n)
		if n.Data == "a" {
			var l Link
			addLink(n, &l)
			*links = append(*links, l)
			return
		}
	// All other cases are here in case the user want to do something with the other nodes...
	// Right now, the nodes are displayed if verbose = true, else nothing is done.
	case html.TextNode:
		displayNode("Text", n)
	case html.DocumentNode:
		displayNode("Document", n)
	case html.CommentNode:
		displayNode("Comment", n)
	case html.DoctypeNode:
		displayNode("Doctype", n)
	case html.RawNode:
		displayNode("Raw", n)
	default:
		log.Printf("Unknown node type: %+v\n", n.Type)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		parse(c, links)
	}
}

// Utility function : display a Node (Text and Attr) (Useful to debug...)
func displayNode(s string, n *html.Node) {
	if verbose {
		fmt.Printf("%s node : %v\n", s, n.Data)
		for _, a := range n.Attr {
			fmt.Printf("\tKey = %v, Val = %v\n", a.Key, a.Val)
		}
	}
}

// Recursive : if node is an Anchor => add it to "l.Href"
//        else if node is a TextNode => add it to "l.Text"
func addLink(n *html.Node, l *Link) {
	for _, a := range n.Attr {
		if a.Key == "href" {
			l.Href = a.Val
			break
		}
	}
	if n.Type == html.TextNode {
		l.Text = l.Text + n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		addLink(c, l)
	}
}

func main() {
	filename := flag.String("f", "ex1.html", "HTML file to parse")
	v := flag.Bool("v", false, "Verbose mode")
	flag.Parse()
	verbose = *v

	bytes, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := html.Parse(strings.NewReader(string(bytes)))
	if err != nil {
		log.Fatal(err)
	}

	// The slice of all anchors found while parsing
	var links []Link

	parse(doc, &links)

	fmt.Printf("%+v\n", links)
}
