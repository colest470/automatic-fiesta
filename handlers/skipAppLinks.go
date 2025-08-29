package handlers

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func SkipAppLinks(link string) bool {
	blockedDormains := []string {"reddit", "youtube", "facebook", "twitter", "instagram", "tiktok", "linkedin", "whatsapp", "discord", "netflix"}

	parsedLink, err := url.Parse(link)
	if err != nil {
		log.Fatalf("Error Parsing url")
	}

	fmt.Println("Base link:", parsedLink.Host)

	dormainHost := strings.Split(parsedLink.Host, ".")
	for _, v := range blockedDormains {
		if v == dormainHost[1] {
			return true
		}
	}
	return false
}