package main

import "github.com/astridalia/wizlib"

func main() {
	checkWiki := wizlib.NewRepository(wizlib.NewHTTPDocumentFetcher(), "https://www.wizard101.com/pvp/pvp-rankings?age=4")
	tours, err := checkWiki.FetchRankings()
	if err != nil {
		panic(err)
	}
	consolePresenter := &wizlib.ConsolePresenter{}
	consolePresenter.PresentRankings(tours)
}
