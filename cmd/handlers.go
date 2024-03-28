package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func registerPageHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err = template.ParseGlob("../templates/*.html")
	if err != nil {
		fmt.Println("Error parsing templates.")
	}

	err := tmpl.ExecuteTemplate(w, "register.html", nil)
	if err != nil {
		fmt.Println("Error executing template:", err)
	}
}


func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	username := r.FormValue("username")
	password := r.FormValue("password")
	
	if userExists(username){
		return;
	}
	
	q, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	// TODO: do something with that error instead of printing it
	if err != nil {
		fmt.Println("Error executing query.");
	}
	defer q.Close()
	
	q.Exec(username, password)
	
	tmpl.ExecuteTemplate(w, "registerAuth.html", nil)
}


func userExists(username string) bool {
	q, err := db.Query("SELECT username FROM users WHERE username = ?", username)
	
	// TODO: do something with that error instead of printing it
	if err != nil {
		fmt.Println("Error querying database.");
	}
	defer q.Close()

	if q.Next() {
		return true
	}

	return true
}
