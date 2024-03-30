package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	uuid "github.com/satori/go.uuid"
)

var users = make(map[string]int)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/me", getUser)
	http.ListenAndServe(":4080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")

	if err != nil {
		id := uuid.NewV4()

		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		})

		generatedNumber := rand.Intn(999999999)
		http.SetCookie(w, &http.Cookie{
			Name:  "user",
			Value: strconv.Itoa(generatedNumber),
		})

		users[id.String()] = generatedNumber
	}

	fmt.Println(cookie)
	io.WriteString(w, "logged in, go to /me")
}

func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "you arent logged in! go to / first.", http.StatusUnauthorized)
		return
	}

	user, err := r.Cookie("user")
	if err != nil {
		http.Error(w, "you arent logged in! go to / first.", http.StatusUnauthorized)
		return
	}

	io.WriteString(w, "Your user is "+id.Value+" and your number is "+user.Value)
}
