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

// ExtractTimestamp extracts the timestamp from a line of text using a regular expression.
// It returns the extracted timestamp and any error encountered.
func ExtractTimestamp(line string) (string, error) {
	re := regexp.MustCompile(`new Date\((\d+)\)`) // Regular expression to match the timestamp

	matches := re.FindStringSubmatch(line) // Find the matches using the regular expression
	if len(matches) < 2 {
		return "", fmt.Errorf("no timestamp found")
	}

	timestamp := matches[1] // Extract the timestamp value from the matches
	return timestamp, nil
}

// FetchHTML fetches the HTML content of a webpage using the provided URL.
// It returns the parsed goquery.Document and any error encountered.
func FetchHTML(url string) (*goquery.Document, error) {
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

// ParseTournaments parses the goquery.Document and extracts the tournament information.
// It returns a slice of Tournament structs and any error encountered.
func ParseTournaments(doc *goquery.Document) ([]Tournament, error) {
	tournaments := make([]Tournament, 0)
	doc.Find(".schedule table tbody tr").Each(func(i int, s *goquery.Selection) {
		tournament := Tournament{
			Name:      strings.TrimSpace(s.Find("td:nth-child(1)").Text()),
			Levels:    strings.TrimSpace(s.Find("td:nth-child(2)").Text()),
			StartTime: strings.TrimSpace(s.Find("td:nth-child(3)").Text()),
			Duration:  strings.TrimSpace(s.Find("td:nth-child(4)").Text()),
		}

		if ts, err := ExtractTimestamp(tournament.StartTime); err == nil {
			tournament.StartTime = ts
		}

		tournaments = append(tournaments, tournament)
	})

	return tournaments, nil
}

// FetchTournaments fetches the tournaments from the specified URL and returns them.
// It returns a slice of Tournament structs and any error encountered.
func FetchTournaments() ([]Tournament, error) {
	doc, err := FetchHTML("https://www.wizard101.com/pvp/schedule")
	if err != nil {
		return nil, err
	}

	tournaments, err := ParseTournaments(doc)
	if err != nil {
		return nil, err
	}

	return tournaments, nil
}

// PrintTournaments prints the details of the provided tournaments.
func PrintTournaments(tournaments []Tournament) {
	for _, t := range tournaments {
		fmt.Printf("Tournament Name: %s\n", t.Name)
		fmt.Printf("Levels: %s\n", t.Levels)
		fmt.Printf("Start Time: %s\n", t.StartTime)
		fmt.Printf("Duration: %s\n", t.Duration)
		fmt.Println("-----------------------------")
	}
}
