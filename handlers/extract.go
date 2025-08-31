package handlers

import (
	// "bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	// "log"
	"net/http"
	// "golang.org/x/net/html"
)


type GoogleSearchResult struct {
	Items []struct {
		Title string `json:"title"`
		Link  string `json:"link"`
	} `json:"items"`
	Error *GoogleAPIError `json:"error,omitempty"`
}

type GoogleAPIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var ValidLinks []string

func ExtractLinks(urlString string) (bool, error) {
	fmt.Println("Fetching...", urlString)
	fmt.Printf("\n")

	resp, err := http.Get(urlString)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var apiError GoogleAPIError
		if err := json.Unmarshal(bodyBytes, &apiError); err == nil && apiError.Message != "" {
			return false, fmt.Errorf("API error: %s (code %d)", apiError.Message, apiError.Code)
		}
		return false, fmt.Errorf("Link not ok! Status: %d", resp.StatusCode)
	}

	// doc, err := html.Parse(bytes.NewReader(bodyBytes))
	// if err != nil {
	// 	return false, fmt.Errorf("Error parsing the html document: %v", err)
	// }

	fmt.Println("Link", urlString, "valid and can be crawled")

	// var buf bytes.Buffer
	// html.Render(&buf, doc)
	// fmt.Println("HTML content:", buf.String())

	// pageContent, err := TraversePageContent(urlString)
	// if err != nil {
	// 	log.Fatalf("Error crawling to search page content")
	// }
	// fmt.Println("This is the search page data:", pageContent)

	var result GoogleSearchResult
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return false, fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Error != nil {
		return false, fmt.Errorf("API error: %s (code %d)", result.Error.Message, result.Error.Code)
	}

	if len(result.Items) == 0 {
		return false, fmt.Errorf("no results found")
	}

	var ResultContent string

	for _, item := range result.Items {
		fmt.Printf("Checking app link: %s\n", item.Link)
		
		if !SkipAppLinks(item.Link) {
			ctx := context.Background()

			fmt.Printf("Found non-app link: %s\n", item.Link)
			fmt.Printf("\n")
			pageContent, err := TraversePageContent(ctx, item.Link)
			if err != nil {
				log.Fatalf("Error reading link")
			}
			ValidLinks = append(ValidLinks, item.Link)

			ResultContent = Crawl(pageContent)

			fmt.Printf("\n")
			fmt.Println("Final result is", ResultContent[:2000])
		} else {
			fmt.Printf("Skipping app link: %s\n", item.Link) // skip then go next link which i cant see in the slice
		}
	}
	fmt.Printf("\n\n")

	fmt.Println("Google result is:", result)
	return true, nil
}