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
