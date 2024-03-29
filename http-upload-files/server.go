package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4050", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	var s string
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("file")
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer f.Close()

		fmt.Println("file: ", f, "\nheader: ", h, "\nerr", err)

		bs, err := io.ReadAll(f)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		s = string(bs)

		dst, err := os.Create("./user/" + h.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err = dst.Write(bs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(
		w,
		"<form method='POST' enctype='multipart/form-data'><input type='file' name='file' /><input type='submit' /></form><br/>"+s,
	)
}
