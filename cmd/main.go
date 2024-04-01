package main

import (
	"net/http"

	"github.com/gorilla/context"
)

var err error

func main() {
	// Database connection
	dbConnect()

	// Register routes
	routes()
	
	// Parse templates
	parseTemplates()

	// Setup static file serving
	setupStaticFileServer()


	// Host the http server
	http.ListenAndServe(":8080", context.ClearHandler(http.DefaultServeMux))	
}