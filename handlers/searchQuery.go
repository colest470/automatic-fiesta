package handlers

import (
	"net/http"
	"time"
	"fmt"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"io"

	"crawler/middlewares"

	"github.com/PuerkitoBio/goquery"
)

func SearchQuery(query string) (string, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return "", fmt.Errorf("error creating cookie jar: %v", err)
	}

	client := &http.Client{
		Timeout: 15 * time.Second,
		Jar:     jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			req.Header = via[len(via)-1].Header
			return nil
		},
	}

	homeReq, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		return "", fmt.Errorf("error creating home request: %v", err)
	}

	middlewares.SetBrowserHeaders(homeReq)
	_, err = client.Do(homeReq)
	if err != nil {
		return "", fmt.Errorf("error fetching google home: %v", err)
	}

	searchURL := "https://www.google.com/search?q=" + url.QueryEscape(query) + "&gl=us&hl=en"

	searchReq, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating search request: %v", err)
	}

	middlewares.SetBrowserHeaders(searchReq)
	searchReq.Header.Set("Referer", "https://www.google.com/")

	time.Sleep(1 * time.Second)

	searchResp, err := client.Do(searchReq)
	if err != nil {
		return "", fmt.Errorf("error fetching search results: %v", err)
	}
	defer searchResp.Body.Close()

	if searchResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code: %d", searchResp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(searchResp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing search results: %v", err)
	}

	var firstResultURL string
	selectors := []string{
		"div.g a",                      // Classic Google
		"div[data-snf] a",              // Modern Google
		"a[jsname=UWckNb]",             // Alternate modern
		"div.yuRUbf a",                 // Another modern variant
		"a[href^='/url?q=']",           // Fallback
	}

	for _, selector := range selectors {
		if firstResultURL != "" {
			break
		}
		doc.Find(selector).EachWithBreak(func(i int, s *goquery.Selection) bool {
			href, exists := s.Attr("href")
			if exists && middlewares.IsValidResultURL(href) {
				firstResultURL = href
				return false // break
			}
			return true // continue
		})
	}

	if firstResultURL == "" {
		return "", fmt.Errorf("no organic results found - check if Google is blocking requests")
	}

	if strings.HasPrefix(firstResultURL, "/url?q=") {
		firstResultURL = strings.Split(firstResultURL, "&")[0][7:]
	}

	resultReq, err := http.NewRequest("GET", firstResultURL, nil)
	if err != nil {
		return "", fmt.Errorf("error creating result page request: %v", err)
	}

	middlewares.SetBrowserHeaders(resultReq)
	resultReq.Header.Set("Referer", searchURL)

	time.Sleep(1 * time.Second)

	resultResp, err := client.Do(resultReq)
	if err != nil {
		return "", fmt.Errorf("error fetching result page: %v", err)
	}
	defer resultResp.Body.Close()

	if resultResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("non-200 status code for result page: %d", resultResp.StatusCode)
	}

	body, err := io.ReadAll(resultResp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading result page: %v", err)
	}

	return string(body), nil
}

// 161429414125-menalsulci622o7guo3c8iivg129jrmp.apps.googleusercontent.com