package main

import (
	"net/http"
)

func routes() {
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginauth", loginAuthHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)

	http.HandleFunc("/addHandler/", addHandler)
	http.HandleFunc("/checkHandler/", checkHandler)
	http.HandleFunc("/uncheckHandler/", uncheckHandler)
	http.HandleFunc("/deleteHandler/", deleteHandler)
	http.HandleFunc("/edit/", editItemHandler)
	http.HandleFunc("/edithandler", editHandler)

}