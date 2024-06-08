package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

	var Db *sql.DB
	var err error

	func DbConnect(){
		username := "root"
		password := ""
		address := "127.0.0.1:3306"
		dbName := "todo"

		dns := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, address, dbName)

		Db, err = sql.Open("mysql", dns)
		if err != nil {
			log.Fatal("Error connecting to database.")
		}
	}