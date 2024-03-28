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

	mux := http.NewServeMux()
	mux.Handle("/dog", dog)
	mux.Handle("/cat", cat)

	http.ListenAndServe(":4080", mux)
}
