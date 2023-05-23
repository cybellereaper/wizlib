package wizlib

import (
	"fmt"
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

// parseFromSelection extracts the player ranking information from a goquery.Selection.
func (r *PlayerRanking) parseFromSelection(s *goquery.Selection) {
	// Extract and store the ranking information from the selection
	r.Position = strings.TrimSpace(s.Find("td:nth-child(1)").Text())
	r.Name = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
	r.Level = strings.TrimSpace(s.Find("td:nth-child(3)").Text())
	r.School = strings.TrimSpace(s.Find("td:nth-child(4) img").AttrOr("class", ""))
	r.Wins = strings.TrimSpace(s.Find("td:nth-child(5)").Text())
	r.Rating = strings.TrimSpace(s.Find("td:nth-child(6)").Text())
}

// FetchRankings retrieves the player rankings from the specified URL.
func FetchRankings(url string) ([]PlayerRanking, error) {
	// Fetch the HTML content from the URL
	doc, err := fetchHTML(url)
	if err != nil {
		return nil, err
	}

	// Parse and return the player rankings from the HTML document
	return parseRankings(doc), nil
}

// parseRankings parses the player rankings from a goquery.Document.
func parseRankings(doc *goquery.Document) []PlayerRanking {
	rankings := make([]PlayerRanking, 0)
	doc.Find(".schedule table tbody tr").Each(func(i int, s *goquery.Selection) {
		// Create a new PlayerRanking instance
		ranking := PlayerRanking{}

		// Parse the ranking information from the selection and store it in the ranking instance
		ranking.parseFromSelection(s)

		// Append the ranking instance to the rankings slice
		rankings = append(rankings, ranking)
	})

	return rankings
}

// PrintRankings prints the player rankings.
func PrintRankings(rankings []PlayerRanking) {
	for _, r := range rankings {
		// Print the ranking information for each player
		fmt.Printf("Position: %s\n", r.Position)
		fmt.Printf("Name: %s\n", r.Name)
		fmt.Printf("Level: %s\n", r.Level)
		fmt.Printf("Wins: %s\n", r.Wins)
		fmt.Printf("School: %s\n", r.School)
		fmt.Printf("Rating: %s\n", r.Rating)
		fmt.Println("-----------------------------")
	}
}
