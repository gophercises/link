package link

import (
	"encoding/json"
	"io"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

type Link struct {
	Href string	`json:"href"`
	Text string	`json:"text"`
}

func (l Link) ToJson() (string, error) {
	j, err :=  json.Marshal(l)

	if err != nil {
		return "", err
	}

	return string(j), nil
}

func Parse(reader io.Reader) (links []Link, err error){
	doc, err := html.Parse(reader)

	if err != nil {
		return
	}

	receiver := make(chan *Link)

	go parseNode(doc, receiver)

	for lp := range receiver {
		links = append(links, *lp)
	}

	return
}

func parseNode(rootNode *html.Node, ch chan<- *Link) {
	defer close(ch)

	wg := sync.WaitGroup{}

	for child := rootNode.FirstChild; child != nil; child = child.NextSibling {
		wg.Add(1)
		go parseNodeRec(child, ch, &wg)
	}

	wg.Wait()
}

func parseNodeRec(node *html.Node, ch chan<- *Link, wg *sync.WaitGroup) {
	defer wg.Done()

	if node.Type == html.ElementNode && node.Data == "a" {
		if link, ok := tryBuildingLink(node); ok {
			ch <- link
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		wg.Add(1)
		go parseNodeRec(child, ch, wg)
	}
}

func tryBuildingLink(node *html.Node) (*Link, bool) {
	var text string
	var href string

	if child := node.FirstChild; child.Type == html.TextNode {
		text = sanitizeText(child.Data)
	}

	if attributes := node.Attr; len(attributes) > 0 {
		for _, attribute := range attributes {
			if attribute.Key == "href" {
				href = attribute.Val
				break
			}
		}
	}

	if len(text) > 0 && len(href) > 0 {
		return &Link{href, text}, true
	}

	return nil, false
}

func sanitizeText(t string) string {
	t = strings.TrimSpace(t)
	t = replaceAllNewLines(t)
	return t
}

var replacement string = " "
var replacer = strings.NewReplacer(
    "\r\n", replacement,
    "\r", replacement,
    "\n", replacement,
    "\v", replacement,
    "\f", replacement,
    "\u0085", replacement,
    "\u2028", replacement,
    "\u2029", replacement,
)
func replaceAllNewLines(t string) string {
	return replacer.Replace(t)
}