package main

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")

		if err != nil {
			id := uuid.NewV4()

			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    id.String(),
				HttpOnly: true,
			})
		}

		fmt.Println(cookie)
	})
	http.ListenAndServe(":4080", nil)
}
