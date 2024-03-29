package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/images.jpeg", serveFile)
	http.ListenAndServe(":4050", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, "<img src='/images.jpeg' />")
}

func serveFile(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("images.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}

	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
	// io.Copy(w, f)

}
