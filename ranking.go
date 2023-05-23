package wizlib

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PlayerRanking struct {
	Position string `json:"position"`
	Name     string `json:"name"`
	Level    string `json:"level"`
	School   string `json:"school"`
	Wins     string `json:"wins"`
	Rating   string `json:"rating"`
}

func (r *PlayerRanking) parseFromSelection(s *goquery.Selection) {
	r.Position = strings.TrimSpace(s.Find("td:nth-child(1)").Text())
	r.Name = strings.TrimSpace(s.Find("td:nth-child(2)").Text())
	r.Level = strings.TrimSpace(s.Find("td:nth-child(3)").Text())
	r.School = strings.TrimSpace(s.Find("td:nth-child(4) img").AttrOr("class", ""))
	r.Wins = strings.TrimSpace(s.Find("td:nth-child(5)").Text())
	r.Rating = strings.TrimSpace(s.Find("td:nth-child(6)").Text())
}

func FetchRankings(url string) ([]PlayerRanking, error) {
	doc, err := FetchHTML(url)
	if err != nil {
		return nil, err
	}
	return ParseRankings(doc), nil
}

func ParseRankings(doc *goquery.Document) []PlayerRanking {
	rankings := make([]PlayerRanking, 0)
	doc.Find(".schedule table tbody tr").Each(func(i int, s *goquery.Selection) {
		ranking := PlayerRanking{}
		ranking.parseFromSelection(s)
		rankings = append(rankings, ranking)
	})
	return rankings
}

func PrintRankings(rankings []PlayerRanking) {
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
