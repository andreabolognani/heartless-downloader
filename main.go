package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func die(rc int, err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(rc)
}

type linkFilterFunc func(string) (bool)

func mp3Filter(s string) (bool) {
	return strings.HasSuffix(s, ".mp3")
}

func extractLinks(url string, filter linkFilterFunc) (links []string) {
	res, err := http.Get(url)
	if err != nil {
		die(2, err)
	}

	doc, err := html.Parse(res.Body)
	if err != nil {
		die(2, err)
	}
	defer res.Body.Close()

	var do func(node *html.Node)
	do = func(node *html.Node) {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, attribute := range node.Attr {
				if atom.Lookup([]byte(attribute.Key)) == atom.Href {
					if (filter(attribute.Val)) {
						links = append(links, attribute.Val)
					}
				}
			}
		}

		for child := node.FirstChild; child != nil; child = child.NextSibling {
			do(child)
		}
	}

	do(doc)
	return
}

func main() {
	if len(os.Args) < 2 {
		die(1, fmt.Errorf("Usage: %s URL", os.Args[0]))
	}

	url := os.Args[1]
	links := extractLinks(url, mp3Filter)

	if len(links) <= 0 {
		die(2, fmt.Errorf("No song found, wrong URL?"))
	} else {
		fmt.Println("Found", len(links), "songs")
		for i := range links {
			fmt.Println(links[i])
		}
	}
}
