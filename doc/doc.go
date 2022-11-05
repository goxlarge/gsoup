package doc

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func GetAttribute(n *html.Node, key string) (string, bool) {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val, true
		}
	}

	return "", false
}

func RenderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	err := html.Render(w, n)
	if err != nil {
		return ""
	}
	return buf.String()
}

func CheckClass(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {
		s, ok := GetAttribute(n, "class")
		if ok && s == className {
			return true
		}
	}

	return false
}

func GetElementsByClass(n *html.Node, className string) *[]*html.Node {
	var targets []*html.Node
	getElementsByClass(n, className, &targets)
	return &targets
}

func getElementsByClass(n *html.Node, className string, targets *[]*html.Node) {
	if CheckClass(n, className) {
		*targets = append(*targets, n)
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getElementsByClass(c, className, targets)
	}
}

func Traverse(n *html.Node) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		fmt.Println(RenderNode(c))
		Traverse(c)
	}
}
