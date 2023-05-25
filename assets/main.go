package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	cloudflarebp "github.com/DaRealFreak/cloudflare-bp-go"
	"github.com/astridalia/wizlib"
)

type WikiText struct {
	Content string `json:"*"`
}

type WikiResponse struct {
	Parse struct {
		Title    string   `json:"title"`
		WikiText WikiText `json:"wikitext"`
	} `json:"parse"`
}

func main() {
	// Specify the API endpoint
	url := "https://www.wizard101central.com/wiki/api.php?action=parse&page=Item:4th_Age_Balance_Talisman&prop=wikitext&formatversion=2&format=json"
	// Create an HTTP client
	client := &http.Client{}
	client.Transport = cloudflarebp.AddCloudFlareByPass(client.Transport)

	// Create a GET request
	req, err := client.Get(url)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return
	}
	defer req.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return
	}

	// Parse the API response
	var api WikiResponse
	if err := json.Unmarshal(body, &api); err != nil {
		fmt.Println("Failed to parse API response:", err)
		return
	}

	info := wizlib.ParsePetInfo(api.Parse.WikiText.Content)

	fmt.Println(info)
}
