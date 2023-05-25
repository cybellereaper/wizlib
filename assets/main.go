package main

import (
	"fmt"

	"github.com/astridalia/wizlib"
)

func main() {
	service := wizlib.NewWikiService()
	content, err := service.GetWikiText("Pet:Bloodbat")
	if err != nil {
		fmt.Println("Failed to fetch wiki text:", err)
		return
	}
	fmt.Println(content)
}
