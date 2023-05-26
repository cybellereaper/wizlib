package main

import (
	"fmt"

	"github.com/astridalia/wizlib"
)

func main() {
	service := wizlib.NewWikiService()
	content, err := service.GetWikiText("Item:4th_Age_Balance_Talisman")
	if err != nil {
		fmt.Println("Failed to fetch wiki text:", err)
		return
	}

	fmt.Println("Wiki Text:", content)

	// Fetching player rankings concurrently

	rankingRepo := wizlib.NewRepository(wizlib.NewHTTPDocumentFetcher(), "https://www.wizard101.com/pvp/pvp-rankings?age=4&levels=1-10&filter=storm")
	rankings, err := rankingRepo.FetchRankings()
	if err != nil {
		fmt.Println("Failed to fetch player rankings:", err)
		return
	}

	consolePresenter := &wizlib.ConsolePresenter{}
	consolePresenter.PresentRankings(rankings)

}
