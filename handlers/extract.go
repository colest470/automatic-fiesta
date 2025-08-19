package handlers

import (
	"fmt"
	"net/http"
	"net/url"
	// "bytes"

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

	// var buf bytes.Buffer
	// html.Render(&buf, doc)
	// fmt.Println(buf.String()) // contets of the html

	baseUrl, err := url.Parse(urlString)
	if err != nil {
		return false, fmt.Errorf("Error parsing the link!", err)
	}

	TraversePageContent(doc, baseUrl)

	ValidLinks = append(ValidLinks, urlString);

	return true, nil
}
