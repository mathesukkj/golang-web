package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("q")
		io.WriteString(w, value)
	})
	http.ListenAndServe(":4050", nil)
}
