package wizlib

import (
	"fmt"
	"net/http"
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

// HTTPDocumentFetcher fetches HTML documents using HTTP.
type HTTPDocumentFetcher struct{}

func (f *HTTPDocumentFetcher) Fetch(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
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
		rankings = append(rankings, ranking)
	})
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

// make an optimized parser

func parseName(name string) string {
	switch name {
	case "LightningName":
		return "Quick Match Tournament"
	case "FireAndIceName":
		return "Fire & Ice Perk Tournament"
	case "OldSchoolName":
		return "Classic Tournament"
	case "AlternatingTurns_PipsAtOnceName":
		return "Turn-Based Tournament"
	case "MythAndStormName":
		return "Myth & Storm Perk Tournament"
	case "LifeAndDeathName":
		return "Life & Death Perk Tournament"
	case "BalanceName":
		return "Balance Perk Tournament"
	default:
		return name
	}
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

	age := u.Query().Get("age")
	levels := u.Query().Get("levels")
	filter := u.Query().Get("filter")

	params := &URLParams{
		Age:    age,
		Levels: levels,
		Filter: filter,
	}

	return params, nil
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
