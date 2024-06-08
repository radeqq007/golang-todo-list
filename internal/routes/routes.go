package routes

import (
	"net/http"

	"todo-list/internal/handlers"
)

func Routes() {
	http.HandleFunc("/", handlers.PageNotFoundHandler)

	http.HandleFunc("/list", handlers.ListHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/loginauth", handlers.LoginAuthHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/registerauth", handlers.RegisterAuthHandler)

	http.HandleFunc("/addHandler/", handlers.AddHandler)
	http.HandleFunc("/checkHandler/", handlers.CheckHandler)
	http.HandleFunc("/uncheckHandler/", handlers.UncheckHandler)
	http.HandleFunc("/deleteHandler/", handlers.DeleteHandler)
	http.HandleFunc("/edit/", handlers.EditItemHandler)
	http.HandleFunc("/edithandler", handlers.EditHandler)

}