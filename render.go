package main

import (
	"encoding/json"
	"flag"
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
	rowHeight = 90
)

func main() {
	flag.Parse()

	fileIn := flag.String("data", "colors.json", "A path to JSON data to read")
	fileTmpl := flag.String("template", "colors.svg.tmpl", "A path to SVG template to use")
	fileOut := flag.String("out", "colors.svg", "Name of the rendered SVG")

	f, err := os.Open(*fileIn)
	if err != nil {
		panic(err)
	}

	var temp Temp
	json.NewDecoder(f).Decode(&temp.Languages)

	for i := range temp.Languages {
		temp.Languages[i].X = rowHeight * i
	}

	temp.Height = len(temp.Languages) * rowHeight

	tmpl, err := template.New(*fileTmpl).ParseFiles(*fileTmpl)
	if err != nil {
		panic(err)
	}

	out, err := os.Create(*fileOut)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(out, temp)
	if err != nil {
		panic(err)
	}
}
