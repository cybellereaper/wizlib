package main

import (
	"fmt"

	"github.com/astridalia/wizlib"
)

func main() {

	names, err := wizlib.GetDefaultNames()

	if err != nil {
		panic(err)
	}

	if name, err := wizlib.CreateName("Alyssa", names); err != nil {
		panic(err)
	} else {
		fmt.Println(name)
	}
}
