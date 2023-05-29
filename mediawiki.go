package wizlib

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type WikiText struct {
	Content string `json:"*"`
}

type WikiResponse struct {
	Parse struct {
		Title    string   `json:"title"`
		Images   []string `json:"images"`
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
func (s *WikiService) GetWikiText(pageID string) (WikiResponse, error) {
	url := fmt.Sprintf("%s?action=parse&page=%s&prop=wikitext|images&formatversion=2&format=json", apiURL, pageID)
	body, err := s.Client.Get(url)
	if err != nil {
		return WikiResponse{}, err
	}

	api, err := s.ParseWikiText(body)
	if err != nil {
		return WikiResponse{}, err
	}

	return *api, nil
}

// ParseWikiText parses the response body into a WikiResponse struct.
func (s *WikiService) ParseWikiText(body []byte) (*WikiResponse, error) {
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

// ParseToJson converts the infobox in the WikiText content to a JSON string.
func (s *WikiService) ParseToJson(pageID string) ([]byte, error) {
	wiki, err := s.GetWikiText(pageID)
	if err != nil {
		return nil, err
	}

	header := FindHeader(wiki.Parse.WikiText.Content)
	infobox := ReplaceInfoboxHeader(wiki.Parse.WikiText.Content, header)
	data := ExtractInfoboxData(infobox)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}

// ReplaceInfoboxHeader removes the infobox header and footer from the WikiText content.
func ReplaceInfoboxHeader(data, template string) string {
	data = strings.TrimPrefix(data, fmt.Sprintf("{{%s\n", template))
	data = strings.TrimSuffix(data, "}}")
	data = strings.TrimSpace(data)
	return data
}

// FindHeader returns the infobox header from the WikiText content.
func FindHeader(data string) string {
	header := regexp.MustCompile(`{{(.+?)\n`)
	headerMatches := header.FindStringSubmatch(data)
	if len(headerMatches) != 2 {
		panic("invalid infobox")
	}
	return headerMatches[1]
}

// ExtractInfoboxData extracts key-value pairs from the infobox.
func ExtractInfoboxData(infobox string) map[string]string {
	re := regexp.MustCompile(`\|([^=]+)=(.*)`)
	matches := re.FindAllStringSubmatch(infobox, -1)

	data := make(map[string]string)

	for _, match := range matches {
		key := strings.TrimSpace(match[1])
		value := strings.TrimSpace(match[2])
		if value != "" {
			data[key] = value
		}
	}

	return data
}
