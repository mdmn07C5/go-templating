package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name   string
	Height int
}

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*.gohtml"))
}

func main() {
	lel := []string{"LMAO", "PEEPEE", "POOPOO"}
	jej := make(map[string]string)
	n := len(lel)
	for i := 0; i < n; i++ {
		jej[lel[i]] = lel[n-i-1]
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template-ranges-maps.gohtml", jej)
	if err != nil {
		log.Fatal(err)
	}

	dudes := []person{
		{Name: "Apu", Height: 456},
		{Name: "Apustaja", Height: 6546},
	}

	err = tpl.ExecuteTemplate(os.Stdout, "template-struct.gohtml", dudes)
	if err != nil {
		log.Fatal(err)
	}
}
