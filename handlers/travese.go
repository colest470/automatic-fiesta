package handlers

import (
	"errors"
	"fmt"
	"net/url"

	"golang.org/x/net/html"
)

type Link struct {
	URL     string
	IsValid bool
	Error   error
}

func TraversePageContent(doc *html.Node, baseURL *url.URL) {
	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					href := a.Val
					link := Link{URL: href}

					parsedURL, parseErr := url.Parse(href)
					if parseErr != nil {
						link.IsValid = false
						link.Error = fmt.Errorf("URL parsing error: %w", parseErr)
						continue
					}

					resolvedURL := baseURL.ResolveReference(parsedURL)
					link.URL = resolvedURL.String()

					if resolvedURL.Scheme == "http" || resolvedURL.Scheme == "https" {
						if resolvedURL.Host != "" {
							link.IsValid = true
						} else {
							link.IsValid = false
							link.Error = errors.New("URL has no host (e.g., mailto:, #anchor)")
							continue
						}
					} else {
						link.IsValid = false
						link.Error = fmt.Errorf("unsupported scheme: %s", resolvedURL.Scheme)
						continue
					}

					if contains(ValidLinks, link.URL) {
						continue
					}

					ValidLinks = append(ValidLinks, link.URL)

					_, err := ExtractLinks(link.URL)
					if err != nil {
						fmt.Printf("Failed to crawl %s: %v\n", link.URL, err)
					}
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}
	}

	traverse(doc)

	fmt.Println("Traversed links:", ValidLinks)
}


func contains(slice []string, item string)  bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func FirstPageContents(url string) { // logs headers only and its contents 

}
