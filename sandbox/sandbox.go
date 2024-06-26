package sandbox

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"lower":   strings.ToLower,
	"flip":    Flip,
	"capital": Capitalize,
}

func Flip(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

type person struct {
	Name   string
	Height int
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func TryPiping() {
	err := tpl.ExecuteTemplate(os.Stdout, "template-piping.gohtml", "PEEPEEPOOPOO")
	if err != nil {
		log.Fatal(err)
	}
}

func TryStructRange() {
	dudes := []person{
		{Name: "Apu", Height: 456},
		{Name: "Apustaja", Height: 6546},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template-struct.gohtml", dudes)
	if err != nil {
		log.Fatal(err)
	}
}

func TryRangeMapFuncs() {
	lel := []string{"LMAO", "PEEPEE", "POOPOO"}
	jej := make(map[string]string)
	n := len(lel)
	for i := 0; i < n; i++ {
		jej[lel[i]] = lel[n-i-1]
	}

	err := tpl.ExecuteTemplate(os.Stdout, "template-ranges-maps-funcs.gohtml", jej)
	if err != nil {
		log.Fatal(err)
	}
}

func TryNesting() {
	err := tpl.ExecuteTemplate(os.Stdout, "template-to-be-nested-into.gohtml", "lelelelel")
	if err != nil {
		log.Fatal(err)
	}
}

func TryNesting2() {
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", "o shid")
	if err != nil {
		log.Fatal(err)
	}
}
