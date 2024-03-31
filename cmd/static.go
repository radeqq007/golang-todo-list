package main

import "net/http"

func setupStaticFileServer() {
	http.Handle("/styles/", http.StripPrefix("/styles/", http.FileServer(http.Dir("../static/styles"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("../static/scripts"))))
}