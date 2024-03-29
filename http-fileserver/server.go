package main

import (
	"io"
	"net/http"
)

func main() {

	//	http.Handle("/", http.FileServer(http.Dir(".")))
	http.Handle(
		"/resources/",
		http.StripPrefix("/resources", http.FileServer(http.Dir("./resources"))),
	)
	http.HandleFunc("/index", index)
	http.ListenAndServe(":4050", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, "<img src='/resources/images.jpeg' />")
}
