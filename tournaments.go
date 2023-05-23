package wizlib

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Tournament struct {
	Name      string `json:"name"`
	Levels    string `json:"levels"`
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
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

// fetchHTML retrieves the HTML content from the specified URL and returns a *goquery.Document.
func fetchHTML(url string) (*goquery.Document, error) {
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

func (t *Tournament) parseFromSelection(s *goquery.Selection) {
	t.Name = strings.TrimSpace(s.Find("td:nth-child(1)").Text())
	t.Levels = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
	t.StartTime = strings.TrimSpace(s.Find("td:nth-child(3)").Text())
	t.Duration = strings.TrimSpace(s.Find("td:nth-child(4)").Text())
	if timestamp, err := extractTimestamp(t.StartTime); err == nil {
		t.StartTime = timestamp
	}
}

// parseTournaments extracts tournament information from the provided *goquery.Document.
func parseTournaments(doc *goquery.Document) ([]Tournament, error) {
	tournaments := make([]Tournament, 0)
	doc.Find(".schedule table tbody tr").Each(func(i int, s *goquery.Selection) {
		tournament := Tournament{}
		tournament.parseFromSelection(s)
		tournaments = append(tournaments, tournament)
	})
	return tournaments, nil
}

// FetchTournaments retrieves the tournaments from the PvP schedule page and returns them as a slice of Tournament.
func FetchTournaments() ([]Tournament, error) {
	doc, err := fetchHTML("https://www.wizard101.com/pvp/schedule")
	if err != nil {
		return nil, err
	}
	return parseTournaments(doc)
}

// PrintTournaments prints the information of the provided tournaments.
func PrintTournaments(tournaments []Tournament) {
	for _, t := range tournaments {
		fmt.Printf("Tournament Name: %s\n", t.Name)
		fmt.Printf("Levels: %s\n", t.Levels)
		fmt.Printf("Start Time: %s\n", t.StartTime)
		fmt.Printf("Duration: %s\n", t.Duration)
		fmt.Println("-----------------------------")
	}
}
