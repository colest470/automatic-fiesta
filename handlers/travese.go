package handlers

// import (
// 	"golang.org/x/net/html"
// )

// type Link struct {
// 	URL     string
// 	IsValid bool
// 	Error   error
// }

// func TraversePageContent() { // not sure what to put
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
		
// 		traverse(doc) // Start the traversal from the root of the document
// 		return foundLinks, nil
// 	}

// }