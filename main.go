package main

import (
	"fmt"
	"hash/fnv"
	"log"
	"os"
	"sync"
	"io"
)

type Queued struct {
	totalQueued int
	number int
	queue []string
	mu sync.Mutex
}

type CrawledSet struct {
	
	mu sync.Mutex
}

var (
	doOnce sync.Once
	existentFile *os.File
	err error
)

func hashUrl(url string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(url))
	return h.Sum64()
}

func insertIntoFile(content string) bool {
	_, err := existentFile.Write([]byte(fmt.Sprintf("Crawled %s, hash - %d\n", content, hashUrl(content))))
	if err != nil {
		log.Printf("Error writing to the file!")
		return false
	}
	return true
}

func main() {
	doOnce.Do(func() {
		existentFile, err = os.OpenFile("crawledWebsites.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to create/open file: %v", err)
		}
		fmt.Println("File opened successfully!")
	})
	defer existentFile.Close()

	insertTruth := insertIntoFile("https://google.com")
	if !insertTruth {
		log.Fatalf("Error inserting into the file")
	}

	// Need to seek to beginning to read what we just wrote
	_, err = existentFile.Seek(0, 0)
	if err != nil {
		log.Fatal("Error seeking file:", err)
	}

	content, err := io.ReadAll(existentFile)
	if err != nil {
		log.Fatal("Error reading file contents:", err)
	}
	fmt.Println(string(content))
}