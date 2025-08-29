package handlers

import (
	"strings"
	"golang.org/x/net/html"
)

func Crawl(pageContent string) string {
	doc, err := html.Parse(strings.NewReader(pageContent))
	if err != nil {
		return "Error parsing HTML: " + err.Error()
	}

	var result strings.Builder
	
	// result.WriteString("=== HEAD CONTENT ===\n")
	// extractHead(doc, &result)
	
	// Extract body content
	result.WriteString("\n=== BODY CONTENT ===\n")
	extractBody(doc, &result)
	
	return result.String()
}

// // extractHead finds and extracts content from the head tag
// func extractHead(n *html.Node, result *strings.Builder) {
// 	if n.Type == html.ElementNode && n.Data == "head" {
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			extractNodeContent(c, result)
// 		}
// 		return
// 	}
	
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		extractHead(c, result)
// 	}
// }

// extractBody finds and extracts content from the body tag
func extractBody(n *html.Node, result *strings.Builder) {
	if n.Type == html.ElementNode && n.Data == "body" {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractNodeContent(c, result)
		}
		return
	}
	
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		extractBody(c, result)
	}
}

// extractNodeContent extracts text content from a node
func extractNodeContent(n *html.Node, result *strings.Builder) {
	switch n.Type {
	case html.TextNode:
		text := strings.TrimSpace(n.Data)
		if text != "" {
			result.WriteString(text)
			result.WriteString("\n")
		}
	case html.ElementNode:
		// Handle specific elements differently if needed
		switch n.Data {
		case "title":
			result.WriteString("Title: ")
		case "meta":
			if name := getAttr(n, "name"); name != "" {
				result.WriteString("Meta[" + name + "]: " + getAttr(n, "content") + "\n")
			}
		case "script", "style":
			// Skip script and style content
			return
		}
		
		// Recursively process child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			extractNodeContent(c, result)
		}
	}
}

// getAttr gets an attribute value from a node
func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}