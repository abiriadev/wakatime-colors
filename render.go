package main

import (
	"encoding/json"
	"os"
	"text/template"
)

type Temp struct {
	Height    int
	Languages []Language
}

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

	var temp Temp
	json.NewDecoder(f).Decode(&temp.Languages)

	for i := range temp.Languages {
		temp.Languages[i].X = 90 * i
	}

	tmpl, err := template.New(fileTmpl).ParseFiles(fileTmpl)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(fileOut)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(out, temp)
	if err != nil {
		panic(err)
	}
}
