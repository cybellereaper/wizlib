# WizLib

[![Go Reference](https://pkg.go.dev/badge/github.com/astridalia/wizlib.svg)](https://pkg.go.dev/github.com/astridalia/wizlib)

WizLib is a Go package that provides utilities for working with wizard names and game data in the magical world of Wizard101.

## Installation

To use WizLib in your Go project, you can simply import it using Go modules:

```shell
go get github.com/astridalia/wizlib@v1.0.2
```

## Features

- **Name Generation**: Generate valid wizard names based on an accepted names list.
- **Game Data Retrieval**: Fetch player rankings and tournament information from the Wizard101 website.
- **Clean Architecture**: Well-organized codebase following clean architecture principles.

## Usage

### Name Generation

```go
package main

import (
	"fmt"
	"github.com/astridalia/wizlib"
)

func main() {
	nameGenerator := wizlib.NewNameGenerator()

	name, err := nameGenerator.GenerateName("Merle Ambrose")
	if err != nil {
		fmt.Println("Failed to generate name:", err)
		return
	}

	fmt.Println("Generated name:", name)
}
```

### Game Data Retrieval

```go
package main

import (
	"fmt"
	"github.com/astridalia/wizlib"
)

func main() {
	rankingRepo := wizlib.NewRankingRepository("https://www.wizard101.com/pvp/pvp-rankings?age=4&levels=1-10&filter=storm")
	tournamentRepo := wizlib.NewTournamentRepository("https://www.wizard101.com/pvp/tournament-rankings?age=4&levels=1-10&filter=death")

	rankings, err := rankingRepo.FetchRankings()
	if err != nil {
		fmt.Println("Failed to fetch player rankings:", err)
		return
	}

	tournaments, err := tournamentRepo.FetchTournaments()
	if err != nil {
		fmt.Println("Failed to fetch tournaments:", err)
		return
	}

	presenter := wizlib.ConsolePresenter{}
	presenter.PresentRankings(rankings)
	presenter.PresentTournaments(tournaments)
}
```

For detailed documentation, refer to the [**GoDoc**](https://pkg.go.dev/github.com/astridalia/wizlib).

