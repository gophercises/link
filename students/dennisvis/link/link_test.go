package link

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func createNode(content string) (n *html.Node) {
	n, err := html.Parse(strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	return
}

func createAnchorNode(href string, text string) *html.Node {
	textNode := createNode(text)
	attrs := make([]html.Attribute, 1)
	attrs[0] = html.Attribute{
		Namespace: "",
		Key:       "href",
		Val:       href,
	}
	return &html.Node{
		Parent:      nil,
		FirstChild:  textNode,
		LastChild:   textNode,
		PrevSibling: nil,
		NextSibling: nil,
		Type:        html.ElementNode,
		DataAtom:    0,
		Data:        "a",
		Namespace:   "",
		Attr:        attrs,
	}
}

func Test_TraverseText_WhenGivenPlainText_ShouldReturnText(t *testing.T) {
	expected := "I am but text"
	node := createNode(expected)

	result := traverseText(node, "")

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_TraverseText_WhenGivenTextInsideSpan_ShouldReturnText(t *testing.T) {
	expected := "I am but text inside a span"
	node := createNode(fmt.Sprintf(`<span>%s</span>`, expected))

	result := traverseText(node, "")

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_TraverseText_WhenGivenTextInsideSpan_AndWithComment_ShouldReturnText(t *testing.T) {
	expected := "I am but text inside a span"
	node := createNode(fmt.Sprintf(`<span>%s<!-- I am a mere comment --></span>`, expected))

	result := traverseText(node, "")

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_TraverseText_WhenGivenTextInsideAnchor_ShouldReturnText(t *testing.T) {
	expected := "I am but text inside an anchor"
	node := createNode(fmt.Sprintf(`<a href="/i-will-only-lead-you-astray">%s</a>`, expected))

	result := traverseText(node, "")

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_TraverseText_WhenGivenDeeplyNestedText_ShouldReturnText(t *testing.T) {
	expected := "I am but deeply nested text"
	node := createNode(fmt.Sprintf(`<main><div><span>%s</span></div></main>`, expected))

	result := traverseText(node, "")

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_TraverseText_WhenGivenMultipleDeeplyNestedTexts_ShouldReturnCombinedText(t *testing.T) {
	expected := "I am but deeply nested text And so am I"
	node := createNode(`
<main>
  <div>
    <span>I am but deeply nested text</span>
  </div>
</main>
<footer>
  <span>And so am I</span>
</footer>
`)

	result := traverseText(node, "")

	if result != expected {
		t.Errorf("Expected [%s] to equal [%s]", result, expected)
	}
}

func Test_NodeToAnchor_WhenGivenPlainTextNode_ShouldReturnAnchor(t *testing.T) {
	expected := Anchor{"", "I am but text"}
	node := createNode(expected.Text)

	result := nodeToAnchor(node)

	if result != expected {
		t.Errorf("Expected [%+v] to equal [%+v]", result, expected)
	}
}

func Test_NodeToAnchor_WhenGivenAnchorNode_ShouldReturnAnchor(t *testing.T) {
	expected := Anchor{"/i-will-only-lead-you-astray", "I am but text"}
	node := createAnchorNode(expected.Href, expected.Text)

	result := nodeToAnchor(node)

	if result != expected {
		t.Errorf("Expected [%+v] to equal [%+v]", result, expected)
	}
}

func traverseAnchorsTest(t *testing.T, exNr int8, expected []Anchor) {
	f, err := os.Open(fmt.Sprintf("../../ex%d.html", exNr))
	if err != nil {
		panic(err)
	}
	node, err := html.Parse(f)
	if err != nil {
		panic(err)
	}

	result := traverseAnchors(node, make([]Anchor, 0))

	if len(result) != len(expected) {
		t.Errorf("Expected result to have length of %d, was %d", len(expected), len(result))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("Expected [%+v] to equal [%+v]", result[0], expected[0])
		}
	}
}

func Test_TraverseAnchors_WhenGivenEx1_ShouldReturnAllAnchors(t *testing.T) {
	expected := make([]Anchor, 1)
	expected[0] = Anchor{"/other-page", "A link to another page"}

	traverseAnchorsTest(t, 1, expected)
}

func Test_TraverseAnchors_WhenGivenEx2_ShouldReturnAllAnchors(t *testing.T) {
	expected := make([]Anchor, 2)
	expected[0] = Anchor{"https://www.twitter.com/joncalhoun", "Check me out on twitter"}
	expected[1] = Anchor{"https://github.com/gophercises", "Gophercises is on Github !"}

	traverseAnchorsTest(t, 2, expected)
}

func Test_TraverseAnchors_WhenGivenEx3_ShouldReturnAllAnchors(t *testing.T) {
	expected := make([]Anchor, 3)
	expected[0] = Anchor{"#", "Login"}
	expected[1] = Anchor{"/lost", "Lost? Need help?"}
	expected[2] = Anchor{"https://twitter.com/marcusolsson", "@marcusolsson"}

	traverseAnchorsTest(t, 3, expected)
}
