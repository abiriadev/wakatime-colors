package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

type Language struct {
	X     int
	Name  string `json:"name"`
	Color string `json:"color"`
}

const (
	fileIn   = "colors.json"
	fileTmpl = "colors.svg.tmpl"
	fileOut  = "colors.svg"
)

func main() {

	f, err := os.Open(fileIn)
	if err != nil {
		panic(err)
	}

	var languages []Language
	json.NewDecoder(f).Decode(&languages)

	for i := range languages {
		languages[i].X = 90 * i
	}

	for _, l := range languages {
		fmt.Println(l.X)
	}

	tmpl, err := template.New(fileTmpl).ParseFiles(fileTmpl)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(fileOut)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(out, languages)
	if err != nil {
		panic(err)
	}
}
