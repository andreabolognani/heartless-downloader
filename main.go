package main

import (
	"fmt"
	"net/http"
	"os"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func die(rc int, err error) {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(rc)
}

func extractLinks(node *html.Node) (res []string) {
	var do func(node *html.Node)
	do = func(node *html.Node) {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attribute := range node.Attr {
				if atom.Lookup([]byte(attribute.Key)) == atom.Href {
					res = append(res, attribute.Val)
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			do(child)
		}
	}
	do(node)
	return
}

func main() {
	if len(os.Args) < 2 {
		die(1, fmt.Errorf("Usage: %s URL", os.Args[0]))
	}

	url := os.Args[1]

	res, err := http.Get(url)
	if err != nil {
		die(2, err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		die(2, err)
	}
	defer res.Body.Close()

	for _, l := range extractLinks(doc) {
		fmt.Println(l)
	}
}
