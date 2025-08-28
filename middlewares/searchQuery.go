package middlewares

import (
	"fmt"
	"net/url"

	"crawler/handlers"
)

func ParseUrl(apiKey, searchEngineID, query string) (string, error) {
	baseURL := "https://www.googleapis.com/customsearch/v1"
	params := url.Values{}
	params.Add("key", apiKey)
	params.Add("cx", searchEngineID)
	params.Add("q", query)
	params.Add("num", "1")

	apiURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	isValid, error := handlers.ExtractLinks(apiURL)
	if(!isValid) {
		return "", error
	}

	return apiURL, nil
}

	// resp, err := http.Get(apiURL)
	// if err != nil {
	// 	return "", fmt.Errorf("API request failed: %v", err)
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	var apiError GoogleAPIError
	// 	if err := json.NewDecoder(resp.Body).Decode(&apiError); err == nil && apiError.Message != "" {
	// 		return "", fmt.Errorf("API error: %s (code %d)", apiError.Message, apiError.Code)
	// 	}
	// 	return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	// }

	// var result GoogleSearchResult
	// if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
	// 	return "", fmt.Errorf("failed to decode response: %v", err)
	// }

	// if result.Error != nil {
	// 	return "", fmt.Errorf("API error: %s (code %d)", result.Error.Message, result.Error.Code)
	// }

	// if len(result.Items) == 0 {
	// 	return "", fmt.Errorf("no results found")
	// }

	// return fetchPageContent(result.Items[0].Link)