package main

import (
	"io"
	"net/http"
)

type handlerDog int

func (m handlerDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "woof")
}

type handlerCat int

func (m handlerCat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "meow")
}

func main() {
	var dog handlerDog
	var cat handlerCat

	// using Mux
	mux := http.NewServeMux()
	mux.Handle("/cat", cat)

	// using DefaultServeMux and Handle
	http.Handle("/dog", dog)

	// using DefaultServeMux and HandleFunc
	http.HandleFunc("/wolf", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "auuuuuuuuuuuu")
	})

	//http.ListenAndServe(":4080", mux)
	http.ListenAndServe(":4080", nil)

}
