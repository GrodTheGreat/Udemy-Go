package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `
CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	passwordHash TEXT NOT NULL
)
`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic(err)
	}

	createEventTable := `
CREATE TABLE IF NOT EXISTS event (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	datetime TEXT NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES user(id)
)
`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		fmt.Println("failed to create events table")
		panic(err)
	}
}
