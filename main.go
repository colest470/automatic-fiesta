package main

import (
	"crawler/middlewares"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	var doOnce sync.Once

	doOnce.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	})

	apiKey := os.Getenv("GOOGLE_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_API_KEY environment variable not set")
	}

	searchEngineID := os.Getenv("GOOGLE_SEARCH_ENGINE_ID")
	if searchEngineID == "" {
		log.Fatal("GOOGLE_SEARCH_ENGINE_ID environment variable not set")
	}

	parsedURL, err := middlewares.ParseUrl(apiKey, searchEngineID, "how to cook pilau")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	//fmt.Println("Successfully fetched contents of first result:")
	//fmt.Println(contents[:min(1000, len(contents))]) // Print first 1000 chars
	fmt.Println("Parsed URL is: ", parsedURL)
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