package main

import (
	"html/template"
	"io"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, "nossa")
	}))

	http.HandleFunc("/dog/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "dog")
	}))

	http.HandleFunc("/me", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "matheus")
	}))

	http.ListenAndServe(":4080", nil)
}
