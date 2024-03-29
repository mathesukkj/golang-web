package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "index")
	})

	http.HandleFunc("/dog/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "dog")
	})

	http.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "matheus")
	})

	http.ListenAndServe(":4080", nil)
}
