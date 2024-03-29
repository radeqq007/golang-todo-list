package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

	var db *sql.DB

	func dbConnect(){
		username := "root"
		password := ""
		address := "127.0.0.1:3306"
		dbName := "todo"

		dns := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, address, dbName)

		db, err = sql.Open("mysql", dns)
		if err != nil {
			log.Fatal("Error connecting to database.")
		}
	}