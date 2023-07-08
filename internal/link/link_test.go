package link

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// arg means name of file and the expected stands for the 'result we expect'
type linkTest struct {
	arg      string
	expected []Link
}

var linkTests = []linkTest{
	linkTest{"../../ex1.html", []Link{{Url: "/other-page", Text: "A link to another page"}}},
	linkTest{"../../ex2.html", []Link{{Url: "https://www.twitter.com/joncalhoun", Text: "Check me out on twitter"}, {Url: "https://github.com/gophercises", Text: "Gophercises is on Github"}}},
	linkTest{"../../ex3.html", []Link{{Url: "#", Text: "Login"}, {Url: "/lost", Text: "Lost? Need help?"}, {Url: "https://twitter.com/marcusolsson", Text: "@marcusolsson"}}},
	linkTest{"../../ex4.html", []Link{{Url: "/dog-cat", Text: "dog cat"}}},
}

func Test_ExtractLinks(t *testing.T) {

	for _, test := range linkTests {
		var links []Link
		p, err := NewParser(test.arg)
		assert.Nil(t, err)

		doc, err := p.Parse()
		assert.Nil(t, err)

		p.ExtractLinks(doc, &links)

		if !assert.Equal(t, links, test.expected) {
			t.Errorf("Output %q not equal to expected %q", links, test.expected)
		}
	}
}
