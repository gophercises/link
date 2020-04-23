package main

import (
	"flag"
	"io/ioutil"

	"github.com/fenriz07/link/helpers"
)


type Link{
	Href string
	Text string
}

func main() {

	html := getHtml()

	helpers.DD(html)

}

func getHtml() string {
	namefile := flag.String("namefile", "ex1.html", "define to html file to parse")

	flag.Parse()

	bshtml, err := ioutil.ReadFile(*namefile)

	if err != nil {
		helpers.Exit(err)
	}

	return string(bshtml)

}
