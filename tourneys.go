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

func ExtractTimestamp(line string) (string, error) {
	// Define the regular expression pattern to match the timestamp
	re := regexp.MustCompile(`new Date\((\d+)\)`)

	// Find the matches using the regular expression
	matches := re.FindStringSubmatch(line)
	if len(matches) < 2 {
		return "", fmt.Errorf("no timestamp found")
	}

	// Extract the timestamp value from the matches
	timestamp := matches[1]
	return timestamp, nil
}

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

func PrintTournaments(tournaments []Tournament) {
	for _, t := range tournaments {
		fmt.Printf("Tournament Name: %s\n", t.Name)
		fmt.Printf("Levels: %s\n", t.Levels)
		fmt.Printf("Start Time: %s\n", t.StartTime)
		fmt.Printf("Duration: %s\n", t.Duration)
		fmt.Println("-----------------------------")
	}
}
