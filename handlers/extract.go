package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/html"
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

	resp, err := http.Get(urlString)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var  apiError GoogleAPIError
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err == nil && apiError.Message != "" {
			return false, fmt.Errorf("API error: %s (code %d)", apiError.Message, apiError.Code)
		}
		return false, fmt.Errorf("Link not ok!", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return false, fmt.Errorf("Error parsing the html document!", err)
	}

	fmt.Println("Link", urlString, "valid and can be crawled")

	var buf bytes.Buffer
	html.Render(&buf, doc)
	fmt.Println(buf.String()) // contets of the html

	var result GoogleSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Error != nil {
		return false, fmt.Errorf("API error: %s (code %d)", result.Error.Message, result.Error.Code)
	}

	if len(result.Items) == 0 {
		return false, fmt.Errorf("no results found")
	}

	fmt.Println("Google result is:", result)

	TraversePageContent(result.Items[0].Link)

	ValidLinks = append(ValidLinks, urlString);

	return true, nil
}
