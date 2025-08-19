package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"io"
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

func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY environment variable not set")
	}

	searchEngineID := os.Getenv("CLIENT_ID")
	if searchEngineID == "" {
		log.Fatal("GOOGLE_SEARCH_ENGINE_ID environment variable not set")
	}

	contents, err := getFirstSearchResult(apiKey, searchEngineID, "how to cook pilau")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Println("Successfully fetched contents of first result:")
	fmt.Println(contents[:min(1000, len(contents))]) // Print first 1000 chars
}

func getFirstSearchResult(apiKey, searchEngineID, query string) (string, error) {
	baseURL := "https://www.googleapis.com/customsearch/v1"
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("cx", searchEngineID)
	params.Add("q", query)
	params.Add("num", "1")

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiError GoogleAPIError
		if err := json.NewDecoder(resp.Body).Decode(&apiError); err == nil && apiError.Message != "" {
			return "", fmt.Errorf("API error: %s (code %d)", apiError.Message, apiError.Code)
		}
		return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var result GoogleSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %v", err)
	}

	if result.Error != nil {
		return "", fmt.Errorf("API error: %s (code %d)", result.Error.Message, result.Error.Code)
	}

	if len(result.Items) == 0 {
		return "", fmt.Errorf("no results found")
	}

	return fetchPageContent(result.Items[0].Link)
}

func fetchPageContent(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch page: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}

	return string(body), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


// package main

// import (
// 	"fmt"
// 	"hash/fnv"
// 	"log"
// 	"os"
// 	"sync"
// 	//"io"

// 	"crawler/handlers"

// 	"github.com/joho/godotenv"
// )

// type Queued struct {
// 	totalQueued int
// 	number int
// 	queue []string
// 	mu sync.Mutex
// }

// type CrawledSet struct {
// 	mu sync.Mutex
// }

// var (
// 	doOnce sync.Once
// 	existentFile *os.File
// 	err error
// 	FirstLink string
// )

// func hashUrl(url string) uint64 {
// 	h := fnv.New64a()
// 	h.Write([]byte(url))
// 	return h.Sum64()
// }

// // func insertIntoFile(content string) bool {
// // 	_, err := existentFile.Write([]byte(fmt.Sprintf("Crawled %s, hash - %d\n", content, hashUrl(content))))
// // 	if err != nil {
// // 		log.Printf("Error writing to the file!")
// // 		return false
// // 	}
// // 	return true
// // }

// func main() {
// 	doOnce.Do(func() {
// 		existentFile, err = os.OpenFile("crawledWebsites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
// 		if err != nil {
// 			log.Fatalf("Failed to create/open file: %v", err)
// 		}
// 		fmt.Println("File opened successfully!")

// 		err := godotenv.Load()
// 		if err != nil {
// 			log.Fatal("Error loading env variables!")
// 		}
// 	})

// 	fmt.Println(os.Getenv("API_KEY"))

// 	defer existentFile.Close()

// 	FirstLink = "https://example.com"

// 	var sendQuery string = "How to cook pilau"

// 	pageContents, err := handlers.SearchQuery(sendQuery)
// 	if err != nil {
// 		log.Fatal("Error in search query ", err)
// 	}

// 	fmt.Println(pageContents[:1000])

// 	// _, err := handlers.ExtractLinks(FirstLink)
// 	// if err != nil {
// 	// 	log.Fatal("Error trying to fetch/extrat link:", FirstLink)
// 	// 	return
// 	// }
// 	// insertTruth := insertIntoFile(FirstLink) after finishing, put the links in a txt file
// 	// if !insertTruth {
// 	// 	log.Fatalf("Error inserting into the file")
// 	// }

// 	// Need to seek to beginning to read what we just wrote
// 	// _, err = existentFile.Seek(0, 0)
// 	// if err != nil {
// 	// 	log.Fatal("Error seeking file:", err)
// 	// }

// 	// content, err := io.ReadAll(existentFile)
// 	// if err != nil {
// 	// 	log.Fatal("Error reading file contents:", err)
// 	// }
// 	// fmt.Println(string(content))
// }