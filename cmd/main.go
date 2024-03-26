package main

import (
	"go-templates/temp"
	"go-templates/templates"
)

func main() {
	rudimentaryTemplate := temp.CreateRudimentaryTemplate("Apu Apustaja")
	templates.HTMLfromString(rudimentaryTemplate, "index", "")
}
