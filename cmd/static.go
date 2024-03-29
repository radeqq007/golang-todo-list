package main

import "net/http"

func setupStaticFileServer() {
	styles := http.FileServer(http.Dir("../static/styles/"))
	scripts := http.FileServer(http.Dir("../static/scripts/"))

	// doesn't work
	// TODO: fix this later
	http.Handle("/styles/", http.StripPrefix("/styles/", styles))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", scripts))
}