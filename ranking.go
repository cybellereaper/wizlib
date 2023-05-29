package wizlib

import (
	"bytes"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// PlayerRanking represents the ranking information of a player.
type PlayerRanking struct {
	Position string `json:"position"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	School   string `json:"school"`
	Wins     string `json:"wins"`
	Rating   string `json:"rating"`
}

// Tournament represents the information of a tournament.
type Tournament struct {
	Name      string `json:"name"`
	Levels    string `json:"levels"`
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
}

// DocumentFetcher retrieves HTML documents from a source.
type DocumentFetcher interface {
	Fetch(url string) (*goquery.Document, error)
}

// HTTPDocumentFetcher provides methods for fetching HTML documents.
type HTTPDocumentFetcher struct {
	Client *APIClient
}

// NewHTTPDocumentFetcher creates a new instance of HTTPDocumentFetcher.
func NewHTTPDocumentFetcher() *HTTPDocumentFetcher {
	return &HTTPDocumentFetcher{
		Client: NewAPIClient(),
	}
}

// Fetch fetches the HTML document from the specified URL.
func (f *HTTPDocumentFetcher) Fetch(url string) (*goquery.Document, error) {
	body, err := f.Client.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// Repository retrieves data from a source.
type Repository struct {
	DocumentFetcher DocumentFetcher
	URL             string
}

// NewRepository creates a new instance of Repository.
func NewRepository(fetcher DocumentFetcher, url string) *Repository {
	return &Repository{
		DocumentFetcher: fetcher,
		URL:             url,
	}
}

// FetchRankings retrieves the player rankings.
func (r *Repository) FetchRankings() ([]PlayerRanking, error) {
	doc, err := r.DocumentFetcher.Fetch(r.URL)
	if err != nil {
		return nil, err
	}

	return parseRankings(doc), nil
}

// FetchTournaments retrieves the tournaments.
func (r *Repository) FetchTournaments() ([]Tournament, error) {
	doc, err := r.DocumentFetcher.Fetch(r.URL)
	if err != nil {
		return nil, err
	}

	return parseTournaments(doc), nil
}

// parseRankings parses the player rankings from a goquery.Document.
func parseRankings(doc *goquery.Document) []PlayerRanking {
	rankings := make([]PlayerRanking, 0)
	doc.Find(".schedule table tbody tr").Each(func(i int, s *goquery.Selection) {
		ranking := PlayerRanking{}
		ranking.parseFromSelection(s)
		if ranking.Name != "" {
			rankings = append(rankings, ranking)
		}
	})

	// Remove empty values
	rankings = rankings[:len(rankings)-1]
	return rankings
}

// parseTournaments parses the tournaments from a goquery.Document.
func parseTournaments(doc *goquery.Document) []Tournament {
	tournaments := make([]Tournament, 0)
	doc.Find(".schedule table tbody tr").Each(func(i int, s *goquery.Selection) {
		tournament := Tournament{}
		tournament.parseFromSelection(s)
		tournaments = append(tournaments, tournament)
	})

	// Remove empty values
	tournaments = tournaments[:len(tournaments)-1]
	return tournaments
}

// parseFromSelection extracts the ranking information from a goquery.Selection.
func (r *PlayerRanking) parseFromSelection(s *goquery.Selection) {
	r.Position = strings.TrimSpace(s.Find("td:nth-child(1)").Text())
	r.Name = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
	r.Level = strings.TrimSpace(s.Find("td:nth-child(3)").Text())
	r.School = strings.TrimSpace(s.Find("td:nth-child(4) img").AttrOr("class", ""))
	r.Wins = strings.TrimSpace(s.Find("td:nth-child(5)").Text())
	r.Rating = strings.TrimSpace(s.Find("td:nth-child(6)").Text())
}

// parseFromSelection extracts the tournament information from a goquery.Selection.
func (t *Tournament) parseFromSelection(s *goquery.Selection) {
	t.Name = parseName(strings.TrimSpace(s.Find("td:nth-child(1)").Text()))
	t.Levels = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
	t.StartTime = strings.TrimSpace(s.Find("td:nth-child(3)").Text())
	t.Duration = strings.TrimSpace(s.Find("td:nth-child(4)").Text())
	if timestamp, err := extractTimestamp(t.StartTime); err == nil {
		t.StartTime = timestamp
	}
}

var nameMap = map[string]string{
	"LightningName":                   "Quick Match Tournament",
	"FireAndIceName":                  "Fire & Ice Perk Tournament",
	"OldSchoolName":                   "Classic Tournament",
	"AlternatingTurns_PipsAtOnceName": "Turn-Based Tournament",
	"MythAndStormName":                "Myth & Storm Perk Tournament",
	"LifeAndDeathName":                "Life & Death Perk Tournament",
	"BalanceName":                     "Balance Perk Tournament",
}

// parseName parses the name of a tournament.
func parseName(name string) string {
	if val, ok := nameMap[name]; ok {
		return val
	}
	return name
}

// extractTimestamp extracts the timestamp from the given line using a regular expression.
func extractTimestamp(line string) (string, error) {
	re := regexp.MustCompile(`new Date\((\d+)\)`)

	matches := re.FindStringSubmatch(line)
	if len(matches) < 2 {
		return "", fmt.Errorf("no timestamp found")
	}

	return matches[1], nil
}

// Presenter is responsible for presenting data.
type Presenter interface {
	PresentRankings(rankings []PlayerRanking)
	PresentTournaments(tournaments []Tournament)
}

// ConsolePresenter presents player rankings and tournaments on the console.
type ConsolePresenter struct{}

// PresentRankings prints the player rankings.
func (p *ConsolePresenter) PresentRankings(rankings []PlayerRanking) {
	for _, r := range rankings {
		fmt.Printf("Position: %s\n", r.Position)
		fmt.Printf("Name: %s\n", r.Name)
		fmt.Printf("Level: %s\n", r.Level)
		fmt.Printf("Wins: %s\n", r.Wins)
		fmt.Printf("School: %s\n", r.School)
		fmt.Printf("Rating: %s\n", r.Rating)
		fmt.Println("-----------------------------")
	}
}

// PresentTournaments prints the tournament information.
func (p *ConsolePresenter) PresentTournaments(tournaments []Tournament) {
	for _, t := range tournaments {
		fmt.Printf("Tournament Name: %s\n", t.Name)
		fmt.Printf("Levels: %s\n", t.Levels)
		fmt.Printf("Start Time: %s\n", t.StartTime)
		fmt.Printf("Duration: %s\n", t.Duration)
		fmt.Println("-----------------------------")
	}
}

// URLParams contains the parameters for a URL.
type URLParams struct {
	Age    string
	Levels string
	Filter string
}

// URLParser is responsible for parsing URL parameters.
type URLParser struct {
	rawURL string
}

// NewURLParser creates a new instance of URLParser.
func NewURLParser(rawURL string) *URLParser {
	return &URLParser{
		rawURL: rawURL,
	}
}

// ParseURL parses the URL and extracts the relevant parameters.
func (p *URLParser) ParseURL() (*URLParams, error) {
	u, err := url.Parse(p.rawURL)

	if err != nil {
		return nil, err
	}

	return &URLParams{
		Age:    u.Query().Get("age"),
		Levels: u.Query().Get("levels"),
		Filter: u.Query().Get("filter"),
	}, nil
}

// URLGenerator is responsible for generating URLs with parameters.
type URLGenerator struct {
	baseURL string
	params  *URLParams
}

// NewURLGenerator creates a new instance of URLGenerator.
func NewURLGenerator(baseURL string) *URLGenerator {
	return &URLGenerator{
		baseURL: baseURL,
	}
}

// WithParams sets the parameters for the URLGenerator.
func (g *URLGenerator) WithParams(params *URLParams) *URLGenerator {
	g.params = params
	return g
}

// GenerateURL generates a URL with the provided parameters.
func (g *URLGenerator) GenerateURL() (string, error) {
	u, err := url.Parse(g.baseURL)

	if err != nil {
		return "", err
	}

	q := u.Query()
	q.Set("age", g.params.Age)
	q.Set("levels", g.params.Levels)
	q.Set("filter", g.params.Filter)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
