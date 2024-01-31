package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Language struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func main() {
	f, err := os.Open("./colors.json")
	if err != nil {
		panic(err)
	}

	var languages []Language
	json.NewDecoder(f).Decode(&languages)

	for _, l := range languages {
		fmt.Println("name: ", l.Name, "color:", l.Color)
	}
}
