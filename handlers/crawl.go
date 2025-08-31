package handlers

import (
	"strings"

	"golang.org/x/net/html"
)

var Doc *html.Node

func Crawl(pageContent string) string {
	var err error
	Doc, err = html.Parse(strings.NewReader(pageContent))
	if err != nil {
		return "Error parsing HTML: " + err.Error()
	}

	var result strings.Builder

	mainContentFound := extractMainContent(&result)

	if !mainContentFound {
		result.Reset() 
		extractBodyContent(Doc, &result)
	}

	return result.String()
}

func extractMainContent(result *strings.Builder) bool {
	mainContentSelectors := []string{
		"id", "content",
		"id", "main",
		"id", "article",
		"role", "main",
		"class", "content",    
		"class", "main", 
		"class", "article",
		"class", "post-content",
	}

	var mainNode *html.Node
	var findMain func(*html.Node)
	findMain = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for i := 0; i < len(mainContentSelectors); i += 2 {
				key, value := mainContentSelectors[i], mainContentSelectors[i+1]
				if attrValue := getAttr(n, key); attrValue != "" {
					if (key == "class" && strings.Contains(attrValue, value)) || attrValue == value {
						mainNode = n
						return
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if mainNode != nil {
				break 
			}
			findMain(c)
		}
	}

	findMain(Doc)
	if mainNode != nil {
		extractTextContent(mainNode, result)
		return true
	}

	result.WriteString("Could not identify a specific main content area.\n")
	return false
}

func extractBodyContent(n *html.Node, result *strings.Builder) {
	if n.Type == html.ElementNode && n.Data == "body" {
		extractTextContent(n, result)
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractBodyContent(c, result)
	}
}

func extractTextContent(n *html.Node, result *strings.Builder) {
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			result.WriteString(text)
			result.WriteString(" ")
		}
	}
	if n.Type == html.ElementNode {
		switch n.Data {
		case "script", "style", "nav", "header", "footer", "aside":
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractTextContent(c, result)
	}
}

func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}