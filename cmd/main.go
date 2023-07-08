package main

import (
	"fmt"
	"github.com/mishankoGO/link/internal/link"
	"log"
)

func main() {
	pagePath := "ex4.html"
	parser, err := link.NewParser(pagePath)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	var links []link.Link
	parser.ExtractLinks(doc, &links)
	if len(links) != 0 {
		fmt.Println(links)
	}
}
