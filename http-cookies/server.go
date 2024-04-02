package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie := http.Cookie{
			Name:  "cookie-test",
			Value: "test",
		}
		http.SetCookie(w, &cookie)
	})

	http.HandleFunc("/cookie-check", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("cookie-test")
		if err != nil {
			http.Error(w, "you dont have the right cookie", http.StatusBadRequest)
			return
		}
		io.WriteString(w, cookie.Name)
	})

	http.HandleFunc("/clear-cookie", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("cookie-test")
		if err != nil {
			io.WriteString(w, "you already cleared that cookie!")
			return
		}

		cookie.MaxAge = -1

		http.SetCookie(w, cookie)
	})

	http.ListenAndServe(":4080", nil)
}
