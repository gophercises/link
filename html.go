s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`

doc, err := html.Parse(strings.NewReader(s))
if err != nil {
	log.Fatal(err)
}
var f func(*html.Node)
f = func(n *html.Node) {

	if n.Type == html.ElementNode && n.Data == "a" {

		fmt.Printf("%v %v %v %v  %v\n", n.Data, n.Type, n.DataAtom, n.Attr, n.Namespace)
	}

	fmt.Printf("%v\n", n)

	if n.Type == html.TextNode {

		fmt.Printf("%v\n", n.Data)

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		f(c)
	}
}
f(doc)