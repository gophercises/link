package main

import (
	"testing"
)

func TestEx1(t *testing.T) {
	actual := GetLinksFromHTML("examples/ex1.html")
	exp := []Link{
		{"/other-page", "A link to another page"},
	}

	if len(actual) != len(exp) {
		t.Errorf("Wrong number of links")
	}
	for i := range actual {
		if actual[i] != exp[i] {
			t.Errorf("Links number %d differ", i)
		}
	}
}

func TestEx2(t *testing.T) {
	actual := GetLinksFromHTML("examples/ex2.html")
	exp := []Link{
		{"https://www.twitter.com/joncalhoun", "Check me out on twitter"},
		{"https://github.com/gophercises", "Gophercises is on Github!"},
	}

	if len(actual) != len(exp) {
		t.Errorf("Wrong number of links")
	}
	for i := range actual {
		if actual[i] != exp[i] {
			t.Errorf("Links number %d differ", i)
		}
	}
}

func TestEx3(t *testing.T) {
	actual := GetLinksFromHTML("examples/ex3.html")
	exp := []Link{
		{"#", "Login"},
		{"/lost", "Lost? Need help?"},
		{"https://twitter.com/marcusolsson", "@marcusolsson"},
	}

	if len(actual) != len(exp) {
		t.Errorf("Wrong number of links")
	}
	for i := range actual {
		if actual[i] != exp[i] {
			t.Errorf("Links number %d differ", i)
		}
	}
}

func TestEx4(t *testing.T) {
	actual := GetLinksFromHTML("examples/ex4.html")
	exp := []Link{
		{"/dog-cat", "dog cat"},
	}

	if len(actual) != len(exp) {
		t.Errorf("Wrong number of links")
	}
	for i := range actual {
		if actual[i] != exp[i] {
			t.Errorf("Links number %d differ", i)
		}
	}
}
