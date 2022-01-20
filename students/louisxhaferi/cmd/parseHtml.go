package main

import (
	"flag"
	"fmt"
	"link"
	"os"
)

var file *string

func init() {
	file = flag.String("file", "", "A html file to read links from")
	flag.Parse()
}

func main() {
	f, err := os.Open(*file)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	links, err := link.Parse(f)

	if err != nil {
		panic(f)
	}
	
	for _, link := range links {
		if j, err := link.ToJson(); err == nil {
			fmt.Println(j)
		}
	}
}
