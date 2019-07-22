package main

import (
	"Manan/link"
	"flag"
	"fmt"
	"os"
)

func main() {
	filename := flag.String("file", "exam.html", "The HTML file to parse links from")
	flag.Parse()
	s, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	links, err := link.Parse(s)
	if err != nil {
		panic(err)
	}

	for _, i := range links {
		fmt.Println("Href: ", i.Href)
		fmt.Println("Text: ", i.Text)
	}
	return
}
