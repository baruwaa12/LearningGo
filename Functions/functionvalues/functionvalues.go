package main

import (
	"fmt"
)

func main() {
	fmt.Println("Running...")
}


func forEachNode(n *html.Node, pre, post func (n *html.Node))  {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}


	func startElement(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
			depth++
		}
	}
	
	func endElement(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}

