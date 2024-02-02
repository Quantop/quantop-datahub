package main

import (
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"quantop.xyz/datahub/utils"
)

func getAPIKey() string {
	key, err := utils.ReadSecretFromFile(".secrets", "MARKETAUX_API_KEY")
	if err != nil {
		log.Fatal("Failed to read MARKETAUX_API_KEY: ", err)
		return ""
	}

	return key
}

func main() {
	baseURL, _ := url.Parse("https://api.marketaux.com")

	baseURL.Path += "v1/news/all"

	// Build the query parameters
	params := url.Values{}
	params.Add("api_token", getAPIKey())
	params.Add("language", "en")
	params.Add("symbols", "AAPL,TSLA")
	params.Add("limit", "50")

	baseURL.RawQuery = params.Encode()
	log.Println("Encoded URL is ", baseURL.String())

	// Create a new GET request
	req, err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
		log.Fatal("Failed to create GET request: ", err)
		return
	}

	// Send the request
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Failed to parse response body: ", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))
}
