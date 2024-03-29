package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()

		io.WriteString(w, params.Get("value"))
	})

	http.ListenAndServe(":4050", nil)
}
