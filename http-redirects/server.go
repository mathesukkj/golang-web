package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://google.com.br", 302)
	})

	http.ListenAndServe(":4080", nil)
}
