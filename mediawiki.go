package wizlib

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"sync"
)

type WikiResponse struct {
	Parse struct {
		Title   string   `json:"title"`
		PageID  int64    `json:"pageid"`
		Images  []string `json:"images"`
		Content string   `json:"wikitext"`
	} `json:"parse"`
}

const apiURL = "https://wiki.wizard101central.com/wiki/api.php"

type WikiService struct {
	Client *APIClient
	cache  sync.Map
}

func NewWikiService(client *APIClient) *WikiService {
	return &WikiService{Client: client}
}

func (s *WikiService) GetWikiText(pageName string) (WikiResponse, error) {
	url := fmt.Sprintf("%s?action=parse&page=%s&prop=wikitext|images&formatversion=2&format=json", apiURL, pageName)

	// Check cache first
	if cachedResponse, ok := s.cache.Load(url); ok {
		return cachedResponse.(WikiResponse), nil
	}

	body, err := s.Client.Get(url)
	if err != nil {
		return WikiResponse{}, err
	}

	var response WikiResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return WikiResponse{}, err
	}

	// Cache the response
	s.cache.Store(url, response)

	return response, nil
}

func (s *WikiService) ParseToJSON(pageName string) ([]byte, error) {
	wiki, err := s.GetWikiText(pageName)
	if err != nil {
		return nil, err
	}

	infoboxContent := extractInfobox(wiki.Parse.Content)
	infoboxData := extractInfoboxData(infoboxContent)

	return json.Marshal(infoboxData)
}

var (
	infoboxStartRegex = regexp.MustCompile(`{{[^{]+`)
	infoboxEndRegex   = regexp.MustCompile(`}}`)
	infoboxDataRegex  = regexp.MustCompile(`\|([^=]+)=([^|}}]+)`)
)

func extractInfobox(wikiText string) string {
	start := infoboxStartRegex.FindStringIndex(wikiText)
	if start == nil {
		return ""
	}

	end := infoboxEndRegex.FindStringIndex(wikiText[start[1]:])
	if end == nil {
		return ""
	}

	return wikiText[start[1] : start[1]+end[0]]
}

func extractInfoboxData(infoboxContent string) map[string]string {
	data := make(map[string]string)

	matches := infoboxDataRegex.FindAllStringSubmatch(infoboxContent, -1)
	for _, match := range matches {
		key := strings.TrimSpace(match[1])
		value := strings.TrimSpace(match[2])
		if value != "" {
			data[key] = value
		}
	}

	return data
}
