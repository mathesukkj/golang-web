package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println(r.Form)
		io.WriteString(w, r.Form.Get("q"))
	})
	http.ListenAndServe(":4050", nil)
}
