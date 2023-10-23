package main

import (
	"link/links"
	"log"
	"strings"
)

var htmlDoc = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
  <div class="sa">
  <a href="/other2-page">A link  2 to another page</a>
  <div>
</body>
</html>
`

func main() {
	r := strings.NewReader(htmlDoc)
	paresedlinks, err := links.Parse(r)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(paresedlinks)
}
