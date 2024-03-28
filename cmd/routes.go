package main

import (
	"net/http"
)

func routes() {
	http.HandleFunc("/register", registerPageHandler)
	http.HandleFunc("/registerauth", registerAuthHandler)
}