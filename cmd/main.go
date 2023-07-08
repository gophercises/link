package main

import (
	"github.com/mishankoGO/link/internal/link"
	"log"
)

func main() {
	pagePath := "ex2.html"
	parser, err := link.NewParser(pagePath)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	parser.PrintDoc(doc)
}
