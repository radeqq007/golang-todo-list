package main

import (
	"net/http"
	"text/template"
)

var tmpl *template.Template

func main() {
	tmpl.ParseGlob("../../templates/*.html")

	// Database connection
	dbConnect()

	// Register routes
	routes()

	// Setup static file serving
	setupStaticFileServer()


	// Host the http server
	http.ListenAndServe(":8080", nil)	
}