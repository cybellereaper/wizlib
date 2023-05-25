package main

import "github.com/astridalia/wizlib"

func main() {
	resp := wizlib.NewRepository(&wizlib.HTTPDocumentFetcher{}, "https://www.wizard101.com/pvp/schedule")
	tour, err := resp.FetchTournaments()
	if err != nil {
		panic(err)
	}
	consolePresenter := &wizlib.ConsolePresenter{}
	consolePresenter.PresentTournaments(tour)
}
