package main

import (
	"bufio"
	"fmt"
	"os"

	"link/links"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	file, err := os.Open("../../ex2.html")
	if err != nil {
		return err
	}
	r := bufio.NewReader(file)

	ls, err := links.GetLinksFromReader(r)
	if err != nil {
		return err
	}
	fmt.Println(ls)

	return nil
}
