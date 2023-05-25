package wizlib

import (
	"encoding/json"
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
	return &APIClient{
		Client: &http.Client{
			Transport: cloudflarebp.AddCloudFlareByPass(http.DefaultTransport),
		},
	}
}

func (c *APIClient) Get(url string) ([]byte, error) {
	log.Print("Fetching: ", url)
	http := &http.Client{}
	http.Transport = cloudflarebp.AddCloudFlareByPass(http.Transport)

	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
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

func (s *WikiService) GetWikiText(pageID string) (interface{}, error) {
	url := fmt.Sprintf("%s?action=parse&page=%s&prop=wikitext&formatversion=2&format=json", apiURL, pageID)
	body, err := s.Client.Get(url)
	if err != nil {
		return make(map[string]string, 0), err
	}
	var api WikiResponse
	if err := json.Unmarshal(body, &api); err != nil {
		return make(map[string]string, 0), err
	}
	return api.Parse.WikiText.Content, nil
}
