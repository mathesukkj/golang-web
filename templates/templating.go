package main

import (
	"log"
	"math"
	"os"
	"strings"
	// "text/template"
	"html/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"uc":   strings.ToUpper,
	"ft":   firstThree,
	"fdbl": double,
	"fsq":  square,
}

func double(i float64) float64 {
	return i * 2
}

func square(i float64) float64 {
	return math.Sqrt(i)
}

func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
}

type Seila struct {
	H1 string
	H2 string
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", Seila{"h1", "h2"})
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl2.gohtml", []string{"h1", "h22 bolsonaro"})
	if err != nil {
		log.Fatal(err)
	}

	err = tpl.ExecuteTemplate(
		os.Stdout,
		"tpl3.gohtml",
		map[string]string{"matheus": "piroca", "<script>alert('mth')</script>": "penis"},
	)
	if err != nil {
		log.Fatal(err)
	}

	//	err = tpl.ExecuteTemplate(os.Stdout, "tpl4.gohtml", time.Now().Format(time.Kitchen))
	err = tpl.ExecuteTemplate(os.Stdout, "tpl4.gohtml", 32.0)
	if err != nil {
		log.Fatal(err)
	}
}
