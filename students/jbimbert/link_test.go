package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func Test_addLink_simpleLink(t *testing.T) {
	s := `<a href="/other-page">A link to another page</a>`
	doc, _ := html.Parse(strings.NewReader(s))
	var l Link
	addLink(doc, &l)
	ResultLink := Link{Href: "/other-page", Text: "A link to another page"}
	if l != ResultLink {
		t.Errorf("Wanted %+v\n Got %+v\n", ResultLink, l)
	}
}

func Test_addLink_LinkWithInnerTags(t *testing.T) {
	s := `<a href="https://github.com/gophercises">Gophercises is on <strong>Github</strong>!</a>`
	doc, _ := html.Parse(strings.NewReader(s))
	var l Link
	addLink(doc, &l)
	ResultLink := Link{Href: "https://github.com/gophercises", Text: "Gophercises is on Github!"}
	if l != ResultLink {
		t.Errorf("Wanted %+v\n Got %+v\n", ResultLink, l)
	}
}

func Test_addLink_LinkWithUnwantedTags(t *testing.T) {
	s := `<a href="https://www.twitter.com/joncalhoun">Check me out on twitter<i class="fa fa-twitter" aria-hidden="true"></i></a>`
	doc, _ := html.Parse(strings.NewReader(s))
	var l Link
	addLink(doc, &l)
	ResultLink := Link{Href: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}
	if l != ResultLink {
		t.Errorf("Wanted %+v\n Got %+v\n", ResultLink, l)
	}
}

func Test_addLink_LinkWithComment(t *testing.T) {
	s := `<a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>`
	doc, _ := html.Parse(strings.NewReader(s))
	var l Link
	addLink(doc, &l)
	ResultLink := Link{Href: "/dog-cat", Text: "dog cat "}
	if l != ResultLink {
		t.Errorf("Wanted %+v\n Got %+v\n", ResultLink, l)
	}
}

func Test_parse(t *testing.T) {
	s := `<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>`
	doc, _ := html.Parse(strings.NewReader(s))
	var l []Link
	parse(doc, &l)
	ResultLink := Link{Href: "/other-page", Text: "A link to another page"}
	if len(l) != 1 {
		t.Errorf("Bad length. Wanted %d, got %d\n", 1, len(l))
	}
	if l[0] != ResultLink {
		t.Errorf("Wanted %+v\n Got %+v\n", ResultLink, l)
	}
}
