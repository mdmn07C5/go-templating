package sandbox

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"lower": strings.ToLower,
	"flip":  Flip,
}

func Flip(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

type person struct {
	Name   string
	Height int
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
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
