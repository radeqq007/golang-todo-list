package main

import (
	"net/http"
)

var err error

func main() {
	// Database connection
	dbConnect()

	// Register routes
	routes()

	// Setup static file serving
	setupStaticFileServer()


	// Host the http server
	http.ListenAndServe(":8080", nil)	
}