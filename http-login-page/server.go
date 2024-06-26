package main

import (
	"html/template"
	"log"
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
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/bar", bar)
	log.Println(http.ListenAndServe(":80", nil))
}

func init() {
	tpl = template.Must(template.New("").ParseGlob("/home/ubuntu/templates/*.gohtml"))
}

func getUser(w http.ResponseWriter, r *http.Request) (User, error) {
	userCookie, err := r.Cookie("sessionId")
	if err != nil {
		return User{}, err
	}

	username := dbSessions[userCookie.Value]
	user := dbUsers[username]
	http.SetCookie(w, userCookie)
	return user, nil
}

func index(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(w, r)
	if err != nil {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", user)
}

func bar(w http.ResponseWriter, r *http.Request) {
	user, err := getUser(w, r)
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

func signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tpl.ExecuteTemplate(w, "signup.gohtml", nil)
	}

	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		foundUser := dbUsers[username]
		hashedPassword := []byte(foundUser.Password)

		if bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)) != nil {
			http.Error(w, "wrong username or password", http.StatusUnauthorized)
		}

		sId := uuid.NewV4()

		c := &http.Cookie{
			Name:   "sessionId",
			Value:  sId.String(),
			MaxAge: 3600,
		}

		http.SetCookie(w, c)
		dbSessions[sId.String()] = username

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.SetCookie(w, &http.Cookie{
			Name:   "sessionId",
			MaxAge: -1,
		})

		http.Redirect(w, r, "/signin", http.StatusFound)
	}
}
