package main

import (
	"io"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/index", index)
	http.ListenAndServe(":4050", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<img src='images.jpeg' />")
}
