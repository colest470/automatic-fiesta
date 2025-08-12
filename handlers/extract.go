package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	"bytes"

	"golang.org/x/net/html"
)

var ValidLinks []string // links in the website as <a href="google.com">Google</a>

func ExtractLinks(urlString string) (bool, error) { // checks if links are valid
	fmt.Println("Fetching", urlString)

	resp, err := http.Get(urlString)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("Link not ok!", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return false, fmt.Errorf("Error parsing the html document!", err)
	}

	fmt.Println("Link", urlString, "valid and can be crawled")

	var buf bytes.Buffer
	html.Render(&buf, doc)
	fmt.Println(buf.String())

	baseUrl, err := url.Parse(urlString)
	if err != nil {
		return false, fmt.Errorf("Error parsing the link!", err)
	}

	fmt.Println(baseUrl)

	ValidLinks = append(ValidLinks, urlString);

	return true, nil
}



// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// 	"net/url" // For URL parsing and joining
// 	"strings" // For string manipulation

// 	"golang.org/x/net/html" // For HTML parsing
// )

// // Link represents a discovered hyperlink.
// type Link struct {
// 	URL     string
// 	IsValid bool // True if it's a valid http/https absolute URL
// 	Error   error // Stores parsing/validation error if any
// }

// // extractLinks fetches a webpage, parses its HTML, and extracts all hyperlinks,
// // returning them as a slice of Link structs with validation status.
// func extractLinks(pageURL string) ([]Link, error) {
// 	fmt.Printf("Fetching: %s\n", pageURL)

// 	resp, err := http.Get(pageURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("error fetching URL %s: %w", pageURL, err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("received non-OK HTTP status for %s: %s", pageURL, resp.Status)
// 	}

// 	doc, err := html.Parse(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing HTML for %s: %w", pageURL, err)
// 	}

// 	var foundLinks []Link
// 	baseURL, err := url.Parse(pageURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("error parsing base URL %s: %w", pageURL, err)
// 	}

// 	// Define a recursive function to traverse the HTML tree
// 	var traverse func(*html.Node)
// 	traverse = func(n *html.Node) {
// 		if n.Type == html.ElementNode && n.Data == "a" {
// 			for _, a := range n.Attr {
// 				if a.Key == "href" {
// 					href := a.Val
// 					link := Link{URL: href}

// 					// 1. Parse the href attribute
// 					parsedURL, parseErr := url.Parse(href)
// 					if parseErr != nil {
// 						link.IsValid = false
// 						link.Error = fmt.Errorf("URL parsing error: %w", parseErr)
// 						foundLinks = append(foundLinks, link)
// 						continue // Skip to next attribute/node if parsing fails
// 					}

// 					// 2. Resolve relative URLs to absolute URLs
// 					resolvedURL := baseURL.ResolveReference(parsedURL)
// 					link.URL = resolvedURL.String() // Update URL to its absolute form

// 					// 3. Validate the resolved URL
// 					// Check for valid scheme (http/https)
// 					if resolvedURL.Scheme == "http" || resolvedURL.Scheme == "https" {
// 						// Further checks: ensure host is not empty, etc.
// 						if resolvedURL.Host != "" {
// 							link.IsValid = true
// 						} else {
// 							link.IsValid = false
// 							link.Error = errors.New("URL has no host (e.g., mailto:, #anchor)")
// 						}
// 					} else {
// 						link.IsValid = false
// 						link.Error = fmt.Errorf("unsupported scheme: %s", resolvedURL.Scheme)
// 					}

// 					foundLinks = append(foundLinks, link)
// 					break // Found href, move to next <a> tag
// 				}
// 			}
// 		}
// 		// Recursively call for child nodes
// 		for c := n.FirstChild; c != nil; c = c.NextSibling {
// 			traverse(c)
// 		}
// 	}

// 	traverse(doc) // Start the traversal from the root of the document
// 	return foundLinks, nil
// }

// func main() {
// 	// Example Usage:
// 	targetURL := "https://en.wikipedia.org/wiki/Go_(programming_language)"
// 	// targetURL := "https://www.example.com"
// 	// targetURL := "http://nonexistent-website-12345.com" // For testing error handling

// 	fmt.Printf("--- Go Web Crawler with Link Validation ---\n")

// 	links, err := extractLinks(targetURL)
// 	if err != nil {
// 		log.Fatalf("Crawler failed: %v\n", err)
// 	}

// 	fmt.Printf("\nSummary of links found on %s:\n", targetURL)
// 	validCount := 0
// 	invalidCount := 0

// 	for i, link := range links {
// 		status := "VALID"
// 		errorMsg := ""
// 		if !link.IsValid {
// 			status = "INVALID"
// 			invalidCount++
// 			if link.Error != nil {
// 				errorMsg = fmt.Sprintf(" (Reason: %v)", link.Error)
// 			}
// 		} else {
// 			validCount++
// 		}
// 		fmt.Printf("%d. %s - %s%s\n", i+1, link.URL, status, errorMsg)
// 	}

// 	fmt.Printf("\n--- Crawl Results ---\n")
// 	fmt.Printf("Total links processed: %d\n", len(links))
// 	fmt.Printf("Valid HTTP/HTTPS links: %d\n", validCount)
// 	fmt.Printf("Invalid/Unsupported links: %d\n", invalidCount)
// 	fmt.Println("----------------------")
// }


