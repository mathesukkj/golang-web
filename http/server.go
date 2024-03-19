package main

import (
	"fmt"
	"net/http"
)

type uai int

func (m uai) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "aaahhahaha")
}

func main() {
	var handler uai
	err := http.ListenAndServe(":5051", handler)
	if err != nil {
		fmt.Println("errorororororor", err)
	}
}
