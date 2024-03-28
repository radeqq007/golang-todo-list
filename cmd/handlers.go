package main

import (
	"fmt"
	"net/http"
)

func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register.html", nil)
}

func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	username := r.FormValue("username")
	//password := r.FormValue("password")

	if userExists(username){
		return;
	}

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
