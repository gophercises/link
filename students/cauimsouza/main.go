package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

// GetLinksFromHTML returns a slice containing information
// about the link tags contained in a HTML file.
// GetLinksFromHTML is NOT thread-safe.
func GetLinksFromHTML(fname string) []Link {
	return getLinks(getParseTreeFromFilename(fname))
}

type Link struct {
	Href string
	Text string
}

func main() {
	printLinks(GetLinksFromHTML(getFilenameFromInput()))
}

func getFilenameFromInput() string {
	fname := flag.String("file", "examples/ex1.html", "HTML file")
	flag.Parse()

	return *fname
}

func getParseTreeFromFilename(fname string) *html.Node {
	file, err := os.Open(fname)
	if err != nil {
		panic(fmt.Sprintf("Error: could not open file %s", fname))
	}

	root, err := html.Parse(file)
	if err != nil {
		panic(fmt.Sprintf("Error: could not parse HTML file"))
	}

	return root
}

func getLinks(root *html.Node) []Link {
	links = nil
	for node := root; node != nil; node = node.NextSibling {
		dfs(node)
	}

	res := make([]Link, len(links))
	copy(res, links)

	return res
}

func dfs(root *html.Node) {
	if isLinkNode(root) {
		links = append(links, getLink(root))
		return
	}

	for child := root.FirstChild; child != nil; child = child.NextSibling {
		dfs(child)
	}
}

func isLinkNode(node *html.Node) bool {
	return node.Type == html.ElementNode && node.Data == "a"
}

func getLink(root *html.Node) Link {
	dest := getLinkDest(root)
	txt := getLinkTxt(root)

	return Link{dest, txt}
}

func getLinkDest(root *html.Node) string {
	for _, attr := range root.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}

	panic("getLinkDest: broken link")
}

func getLinkTxt(root *html.Node) string {
	var s string

	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Type == html.TextNode {
			s = concatenateStrings(s, node.Data)
		}

		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(root)

	return strings.TrimSpace(s)
}

func concatenateStrings(s, t string) string {
	t = strings.TrimSpace(t)
	if len(s) > 0 {
		return s + " " + t
	}

	return t
}

func printLinks(links []Link) {
	fmt.Printf("%d links\n", len(links))

	for i, link := range links {
		fmt.Printf("Link #%d:\n\tdest: %s\n\ttext: %s\n\n", i+1, link.Href, link.Text)
	}
}

var links []Link
