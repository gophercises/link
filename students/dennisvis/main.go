package link

import (
	"flag"
	"fmt"
	"os"

	"github.com/DennisVis/link/students/dennisvis/link"
)

var htmlFile = flag.String("htmlFile", "", "Thw HTML file to parse for links")

func main() {
	flag.Parse()
	if *htmlFile == "" {
		panic("A valid HTML file name is required")
	}

	f, err := os.Open(*htmlFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	anchors, err := link.ParseAnchors(f)
	if err != nil {
		panic(err)
	}

	formatLine := "----------------------------------------"
	fmt.Println(formatLine)
	for _, anchor := range anchors {
		fmt.Printf("href: %s\ntext: %s\n", anchor.Href, anchor.Text)
		fmt.Println(formatLine)
	}
}
