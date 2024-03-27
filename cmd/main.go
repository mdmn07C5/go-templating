package main

import (
	"go-templates/temp"
	"go-templates/templates"
	"log"
	"os"
	"text/template"
)

func main() {
	rudimentaryTemplate := temp.CreateRudimentaryTemplate("Apu Apustaja")
	templates.HTMLfromString(rudimentaryTemplate, "index", "")

	tpl, err := template.ParseGlob("../templates/*.gohtml")
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "template.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
