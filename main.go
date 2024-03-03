package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define a structure to unmarshal the JSON data
type CryptoRates struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
}

func main() {
	// CoinGecko API URL for current Bitcoin price in USD
	url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

	// Make an HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Error retrieving data: ", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body: ", err)
	}

	// Unmarshal the JSON data into the CryptoRates struct
	var rates CryptoRates
	if err := json.Unmarshal(body, &rates); err != nil {
		log.Fatal("Error unmarshaling JSON: ", err)
	}

	// Print the Bitcoin price in USD
	fmt.Printf("Current Bitcoin price in USD: $%.2f\n", rates.Bitcoin.Usd)
}
