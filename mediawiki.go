package wizlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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

type APIClient struct {
	Client *http.Client
}

func NewAPIClient() *APIClient {
	client := &http.Client{}
	client.Transport = cloudflarebp.AddCloudFlareByPass(client.Transport)
	return &APIClient{
		Client: client,
	}
}

func (c *APIClient) Get(url string) ([]byte, error) {
	log.Print("Fetching: ", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
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
