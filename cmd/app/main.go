package main

import (
	"fmt"
	"log"
	"net/http"

	"todo-list/internal/database"
	"todo-list/internal/handlers"
	"todo-list/internal/routes"
	"todo-list/internal/static"

	"github.com/gorilla/context"
)

var err error

func main() {
	// Database connection
	database.DbConnect()

	// Register routes
	routes.Routes()
	
	// Parse templates
	handlers.ParseTemplates()

	// Setup static file serving
	static.SetupStaticFileServer()


	// Host the http server
	port := ":8080"
	fmt.Printf("🚀 Running todo list on localhost%s...\n", port)
	err = http.ListenAndServe(port, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatal(err)
	}

}