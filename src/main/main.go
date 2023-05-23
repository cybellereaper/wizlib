package main

import (
	"github.com/astridalia/wizlib"
)

func main() {
	rep := wizlib.NewRepository(&wizlib.HTTPDocumentFetcher{}, "https://www.wizard101.com/pvp/pvp-rankings?age=4")
	t, err := rep.FetchRankings()
	if err != nil {
		panic(err)
	}
	consolePresenter := &wizlib.ConsolePresenter{}
	consolePresenter.PresentRankings(t)

}
