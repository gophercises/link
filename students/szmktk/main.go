package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func (l Link) String() string {
	return fmt.Sprintf("{href='%s', text='%s'}", l.Href, l.Text)
}

func main() {
	inputFile := readFile(*parseUserInput())
	defer inputFile.Close()
	links := parseLinks(*inputFile)
	log.Println(links)
}

func parseUserInput() *string {
	htmlFilePath := flag.String("file", "", "Path to the HTML file")
	flag.Parse()

	if *htmlFilePath == "" {
		flag.Usage()
		log.Fatalln("Error: HTML file path is required.")
	}

	return htmlFilePath
}

func readFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error reading HTML file: '%s': %v", path, err)
	}

	return file
}

func parseLinks(file os.File) []Link {
	tokenizer := html.NewTokenizer(&file)
	var links []Link
	var buffer bytes.Buffer
	var catchText bool
	var link Link

	for {
		tokenType := tokenizer.Next()
		err := processErrorToken(tokenizer, tokenType)
		if err != nil {
			break
		}

		switch tokenType {
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.DataAtom.String() == "a" && len(token.Attr) > 0 {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link.Href = attr.Val
						catchText = true
					}
				}
			}

		case html.TextToken:
			if catchText {
				buffer.Write(tokenizer.Raw())
			}

		case html.EndTagToken:
			token := tokenizer.Token()
			if token.DataAtom.String() == "a" {
				link.Text = strings.TrimSpace(buffer.String())
				links = append(links, link)
				buffer.Reset()
				catchText = false
			}
		}
	}

	return links
}

func processErrorToken(tokenizer *html.Tokenizer, tokenType html.TokenType) error {
	if tokenType == html.ErrorToken {
		err := tokenizer.Err()
		if err != io.EOF {
			log.Fatalln("Error when parsing HTML", err)
		}
		return err
	}
	return nil
}

