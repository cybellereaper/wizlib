package wizlib

import (
	"encoding/json"
	"errors"
	"fmt"
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

// WikiService provides methods for interacting with the Wizard101 Central Wiki.
type WikiService struct {
	Client *APIClient
}

// NewWikiService creates a new instance of WikiService.
func NewWikiService() *WikiService {
	return &WikiService{
		Client: NewAPIClient(),
	}
}

// GetWikiText returns the WikiText content of the specified page.
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

// parseWikiResponse parses the response body into a WikiResponse struct.
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
