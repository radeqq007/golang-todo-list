package main

import (
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

var tmpl *template.Template

func parseTemplates(){
	tmpl, err = template.ParseGlob("../templates/*.html")
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register.html", nil)
}


func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	username := r.FormValue("username")
	password := r.FormValue("password")
	
	if userExists(username){
		tmpl.ExecuteTemplate(w, "registerauth.html", "Account already exists.")
		return;
	}

	hash, err := hashPassword(password)
	// I should probably do something with this error but I don't know what
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}
	
	q, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	// Same with this error
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()
	
	q.Exec(username, hash)
	
	tmpl.ExecuteTemplate(w, "registerauth.html", nil)
}


func userExists(username string) bool {
	q := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username)
	
	var count int
	q.Scan(&count)

	return count > 0
}

func hashPassword(password string) (string, error)	 {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(byte), err
}


func loginHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
	username := r.FormValue("username")
	password := r.FormValue("password")

	var hash string;

	q := db.QueryRow("SELECT password FROM users WHERE username = ?", username)
	err = q.Scan(&hash)
	if err != nil {
		tmpl.ExecuteTemplate(w, "login.html", "Check username and password.")
		return;
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "login.html", "Check username and password.")
}



func listHandler(w http.ResponseWriter, r *http.Request){
	tmpl.ExecuteTemplate(w, "list.html", nil)
}