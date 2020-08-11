package links

import (
	"bufio"
	"os"
	"testing"
)

func getLinksFromFile(fileName string) ([]Link, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(file)

	return GetLinksFromReader(r)
}

func TestGetSimpleLink(t *testing.T) {
	links, err := getLinksFromFile("../../../ex1.html")
	if err != nil {
		t.Error(err)
	}

	wantLinks := []Link{{Href: "/other-page", Text: "A link to another page"}}

	for i, got := range links {
		want := wantLinks[i]

		if got != want {
			t.Errorf("links are not equal: got: %+v, want: %+v", got, want)
		}
	}
}

func TestGetLinksWithNestedText(t *testing.T) {
	links, err := getLinksFromFile("../../../ex2.html")
	if err != nil {
		t.Error(err)
	}

	wantLinks := []Link{
		{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
		{
			Href: "https://github.com/gophercises",
			Text: "Gophercises is on Github!",
		},
	}

	for i, got := range links {
		want := wantLinks[i]

		if got != want {
			t.Errorf("links are not equal: got: %+v, want: %+v", got, want)
		}
	}
}

func TestGetLinksFromRealDocument(t *testing.T) {
	links, err := getLinksFromFile("../../../ex3.html")
	if err != nil {
		t.Error(err)
	}

	wantLinks := []Link{
		{
			Href: "#",
			Text: "Login",
		},
		{
			Href: "/lost",
			Text: "Lost? Need help?",
		},
		{
			Href: "https://twitter.com/marcusolsson",
			Text: "@marcusolsson",
		},
	}

	for i, got := range links {
		want := wantLinks[i]

		if got != want {
			t.Errorf("links are not equal: got: %+v, want: %+v", got, want)
		}
	}
}

func TestGetLinksWithComments(t *testing.T) {
	links, err := getLinksFromFile("../../../ex4.html")
	if err != nil {
		t.Error(err)
	}

	wantLinks := []Link{
		{
			Href: "/dog-cat",
			Text: "dog cat",
		},
	}

	for i, got := range links {
		want := wantLinks[i]

		if got != want {
			t.Errorf("links are not equal: got: %+v, want: %+v", got, want)
		}
	}
}
