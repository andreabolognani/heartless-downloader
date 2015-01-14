package main

import (
	"fmt"
	"net/http"
	"os"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "URL")
		os.Exit(1)
	}

	url := os.Args[1]

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	var printHrefs func(*html.Node)
	printHrefs = func(node *html.Node) {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attribute := range node.Attr {
				if atom.Lookup([]byte(attribute.Key)) == atom.Href {
					fmt.Println(attribute.Val)
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			printHrefs(child)
		}
	}

	printHrefs(doc)
}
