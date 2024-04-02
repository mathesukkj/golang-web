package main

import (
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	First    string
	UserName string
	Last     string
	Password string
}

var dbUsers = make(map[string]User)
var dbSessions = make(map[string]string)
var userId string
var tpl *template.Template

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":4080", nil)
}

func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}

func getUser(r *http.Request) (User, error) {
	userCookie, err := r.Cookie("sessionId")
	if err != nil {
		return User{}, err
	}

	username := dbSessions[userCookie.Value]
	user := dbUsers[username]
	return user, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(r)
	if err != nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "bar.gohtml", user)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")

		if _, ok := dbUsers[username]; ok {
			http.Error(w, "username taken", http.StatusForbidden)
			return
		}

		sId := uuid.NewV4()

		c := &http.Cookie{
			Name:  "sessionId",
			Value: sId.String(),
		}

		http.SetCookie(w, c)
		dbSessions[sId.String()] = username

		encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 2)
		if err != nil {
			http.Error(w, "some error happened! sorry", http.StatusInternalServerError)
		}

		dbUsers[username] = User{
			First:    firstname,
			UserName: username,
			Last:     lastname,
			Password: string(encryptedPassword),
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
