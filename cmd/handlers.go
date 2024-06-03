package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var tmpl *template.Template
var store = sessions.NewCookieStore([]byte("something-very-secret"))

type ListItem struct {
	ID     		int
	Content 	string
	Checked 	bool
}


func parseTemplates(){
	tmpl, err = template.ParseGlob("../templates/*.html")
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
}


func userExists(username string) bool {
	q := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username)
	
	var count int
	q.Scan(&count)

	return count > 0
}

func hashPassword(password string) (string, error)	 {
	byte, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(byte), err
}

func isLoggedIn(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	_, ok := session.Values["userID"]
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
}

func pageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "404.html", nil)
}


func registerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register.html", nil)
}


func registerAuthHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	username := r.FormValue("username")
	password := r.FormValue("password")
	
	if userExists(username){
		tmpl.ExecuteTemplate(w, "register.html", "User already exists.")
		return;
	}

	hash, err := hashPassword(password)
	if err != nil {
		log.Fatal("Error hashing password:", err)
	}
	
	q, err := db.Prepare("INSERT INTO users (username, password) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()
	
	q.Exec(username, hash)
	
	tmpl.ExecuteTemplate(w, "registerauth.html", nil)
}


func loginHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	_, ok := session.Values["userID"]
	if ok {
		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}

	tmpl.ExecuteTemplate(w, "login.html", nil)
}

func loginAuthHandler(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	
	username := r.FormValue("username")
	password := r.FormValue("password")

	var userID int;
	var hash string;

	q := db.QueryRow("SELECT id, password FROM users WHERE username = ?", username)
	err = q.Scan(&userID, &hash)
	if err != nil {
		tmpl.ExecuteTemplate(w, "login.html", "Check username and password.")
		return;
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == nil {
		session, _ := store.Get(r, "session")
		session.Values["userID"] = userID
		err := session.Save(r, w)
		if err != nil {
			log.Fatal("Error saving session:", err)
			return
		}

		http.Redirect(w, r, "/list", http.StatusSeeOther)
		return
	}
	tmpl.ExecuteTemplate(w, "login.html", "Check username and password.")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["userID"] = nil
	session.Options.MaxAge = -1 
	err := session.Save(r, w)
	if err != nil {
		log.Fatal("Error saving session:", err)
		return
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}


func listHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")

	isLoggedIn(w, r)

	userID := session.Values["userID"]

	rows, err := db.Query("SELECT id, content, checked FROM todos WHERE user_id = ?", userID)
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()

	var items []ListItem

	for rows.Next() {
		var item ListItem
		rows.Scan(&item.ID, &item.Content, &item.Checked)

		items = append(items, item)
	}

	tmpl.ExecuteTemplate(w, "list.html", items)
}

func addHandler(w http.ResponseWriter, r *http.Request) {

	session, _ := store.Get(r, "session")


	isLoggedIn(w, r)

	err = r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form:", err)
	}

	content := r.FormValue("content")
	userID := session.Values["userID"]


	q, err := db.Prepare("INSERT INTO todos (content, checked ,user_id) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()

	q.Exec(content, 0, userID)

	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func uncheckHandler(w http.ResponseWriter, r *http.Request){

	session, _ := store.Get(r, "session")
	isLoggedIn(w, r)

	err = r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form:", err)
	}

	userID := session.Values["userID"]
	id := r.FormValue("id")
	q, err := db.Prepare("UPDATE todos SET checked = 0 WHERE id = ? AND user_id = ?")
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()
	q.Exec(id, userID)
	http.Redirect(w, r, "/list", http.StatusSeeOther)

}

func checkHandler(w http.ResponseWriter, r *http.Request){

	session, _ := store.Get(r, "session")
	isLoggedIn(w, r)

	err = r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form:", err)
	}

	userID := session.Values["userID"]
	id := r.FormValue("id")
	q, err := db.Prepare("UPDATE todos SET checked = 1 WHERE id = ? AND user_id = ?")
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()
	q.Exec(id, userID)
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "session")
	isLoggedIn(w, r)

	err = r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form:", err)
	}

	userID := session.Values["userID"]
	id := r.FormValue("id")
	q, err := db.Prepare("DELETE FROM todos WHERE id = ? AND user_id = ?")
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()
	q.Exec(id, userID)
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}

func editItemHandler(w http.ResponseWriter, r *http.Request){

	session, _ := store.Get(r, "session")
	isLoggedIn(w, r)
	err = r.ParseForm()
	if err != nil {
		log.Fatal("Error parsing form:", err)
	}


	userID := session.Values["userID"]
	var item ListItem
	idStr := r.FormValue("id")
	id, _ := strconv.Atoi(idStr)
	item.ID = id
	q := db.QueryRow("SELECT content FROM todos WHERE id = ? AND user_id = ?", id, userID)
	
	q.Scan(&item.Content)


	tmpl.ExecuteTemplate(w, "edit.html", item)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.FormValue("id")
	content := r.FormValue("content")
	q, err := db.Prepare("UPDATE todos SET content = ? WHERE id = ?")
	if err != nil {
		log.Fatal("Error preparing query:", err)
	}
	defer q.Close()
	q.Exec(content, id)
	http.Redirect(w, r, "/list", http.StatusSeeOther)
}