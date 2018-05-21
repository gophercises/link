package main

import (
	"fmt"
	"github.com/gophercises/link/students/ccallergard"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("plz html path plz")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	ls, err := link.Parse(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(link.LinksString(ls))
}
