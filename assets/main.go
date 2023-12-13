package main

import (
	"fmt"

	"github.com/astridalia/wizlib"
)

func main() {

	checkWiki, err := wizlib.NewWikiService().GetWikiText("Pet:Stormzilla")
	if err != nil {
		panic(err)
	}

	fmt.Println(checkWiki)
	checkWiki2 := wizlib.NewRepository(wizlib.NewHTTPDocumentFetcher(), "https://www.wizard101.com/pvp/pvp-rankings?age=4")
	tours, err := checkWiki2.FetchRankings()
	if err != nil {
		panic(err)
	}
	consolePresenter := &wizlib.ConsolePresenter{}
	consolePresenter.PresentRankings(tours)
}
