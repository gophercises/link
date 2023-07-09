package main

import (
	"fmt"
	"github.com/mishankoGO/link/internal/link"
	"log"
)

func main() {
	pagePath := "ex3.html"
	parser, err := link.NewParser(pagePath)
	if err != nil {
		log.Fatal(err)
	}

	extractLinks, err := parser.ExtractLinks()
	if err != nil {
		log.Fatal(err)
	}
	if len(extractLinks) != 0 {
		fmt.Println(extractLinks)
	}
}
