package middlewares

import (
	"strings"
)

func IsValidResultURL(href string) bool { // for queries
	return href != "" &&
		!strings.HasPrefix(href, "/search?") &&
		!strings.Contains(href, "google.com") &&
		!strings.HasPrefix(href, "#") &&
		!strings.HasPrefix(href, "/images") &&
		!strings.HasPrefix(href, "/maps") &&
		!strings.HasPrefix(href, "/preferences") &&
		!strings.HasPrefix(href, "/advanced_search")
}