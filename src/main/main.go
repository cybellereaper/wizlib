package main

import (
	"github.com/astridalia/wizlib"
)

func main() {
	url := wizlib.NewURLGenerator("https://www.wizard101.com/pvp/pvp-rankings")
	url.WithParams(&wizlib.URLParams{Filter: "storm", Age: "4"})
	urlParsed, err := url.GenerateURL()
	if err != nil {
		panic(err)
	}
	rep := wizlib.NewRepository(&wizlib.HTTPDocumentFetcher{}, urlParsed)
	t, err := rep.FetchRankings()
	if err != nil {
		panic(err)
	}
	consolePresenter := &wizlib.ConsolePresenter{}
	consolePresenter.PresentRankings(t)
}
