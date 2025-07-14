package main

import (
	"fmt"
	"hash/fnv"
	"sync"
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

func hashUrl(url string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(url))
	return  h.Sum64()
}

func main() {
	fmt.Println(hashUrl("https://google.com"))
}