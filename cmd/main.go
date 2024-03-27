package main

import (
	"go-templates/temp"
	"go-templates/templates"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("../templates/*.gohtml"))
}

func main() {
	rudimentaryTemplate := temp.CreateRudimentaryTemplate("Apu Apustaja")
	templates.HTMLfromString(rudimentaryTemplate, "index", "")

	err := tpl.ExecuteTemplate(os.Stdout, "template2.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "template.gohtml", nil)
	if err != nil {
		log.Fatal(err)
	}
}
