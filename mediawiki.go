package wizlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	cloudflarebp "github.com/DaRealFreak/cloudflare-bp-go"
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

const (
	apiURL = "https://www.wizard101central.com/wiki/api.php"
)

// APIClient provides methods for making HTTP requests.
type APIClient struct {
	Client *http.Client
}

// NewAPIClient creates a new instance of APIClient.
func NewAPIClient() *APIClient {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	client.Transport = cloudflarebp.AddCloudFlareByPass(client.Transport)
	return &APIClient{
		Client: client,
	}
}

// Get sends a GET request to the specified URL and returns the response body.
func (c *APIClient) Get(url string) ([]byte, error) {
	defer func() {
		start := time.Now()
		elapsed := time.Since(start)
		log.Printf("Execution time: %s, URL: %s", elapsed, url)
	}()
	// Create a channel to communicate the result
	resultChan := make(chan []byte, 1)
	errChan := make(chan error, 1)

	go func() {
		resp, err := c.Client.Get(url)
		if err != nil {
			errChan <- err
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			errChan <- err
			return
		}

		resultChan <- body
	}()

	// Wait for the goroutine to finish or return an error
	select {
	case body := <-resultChan:
		return body, nil
	case err := <-errChan:
		return nil, err
	}
}

type WikiService struct {
	Client *APIClient
}

func NewWikiService() *WikiService {
	return &WikiService{
		Client: NewAPIClient(),
	}
}

func (s *WikiService) GetWikiText(pageID string) (string, error) {
	url := fmt.Sprintf("%s?action=parse&page=%s&prop=wikitext&formatversion=2&format=json", apiURL, pageID)
	body, err := s.Client.Get(url)
	if err != nil {
		return "", err
	}

	api, err := s.parseWikiResponse(body)
	if err != nil {
		return "", err
	}

	return api.Parse.WikiText.Content, nil
}

func (s *WikiService) parseWikiResponse(body []byte) (*WikiResponse, error) {
	var api WikiResponse
	err := json.Unmarshal(body, &api)
	if err != nil {
		return nil, err
	}

	if api.Parse.WikiText.Content == "" {
		return nil, errors.New("empty WikiText content")
	}

	return &api, nil
}
