# WizLib

[![Go Reference](https://pkg.go.dev/badge/github.com/astridalia/wizlib.svg)](https://pkg.go.dev/github.com/astridalia/wizlib)

WizLib is a Go package that provides utilities for working with wizard names and game data in the magical world of Wizard101.

## Installation

To use WizLib in your Go project, you can simply import it using Go modules:

```shell
go get github.com/astridalia/wizlib
```

## Features

- **Mediawiki**: Get information from the wiki.
- **Name Generation**: Generate valid wizard names based on an accepted names list.
- **Clean Architecture**: Well-organized codebase following clean architecture principles.

## Usage

### Mediawiki Retrieval

```go
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
	fmt.Println(content)
}
```

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

For detailed documentation, refer to the [**GoDoc**](https://pkg.go.dev/github.com/astridalia/wizlib).

