package main

import (
	"database/sql"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"mahi/handlers"
)

func main() {
	// Create database object
	db := initDB("storage.db")
	migrate(db)

	// Create a new instance of Echo
	e := echo.New()

	// Create all routes

	// Static routes for main page and manage page
	e.File("/", "public/index.html")
	e.File("/manage", "public/manage.html")

	// Routes for persons
	e.GET("/persons", handlers.GetPersons(db))
	e.PUT("/persons", handlers.PutPerson(db))
	e.DELETE("/persons/:id", handlers.DeletePerson(db))

	// Routes for attributes
	e.GET("/attributes", handlers.GetAttributes(db))
	e.PUT("/attributes", handlers.PutAttribute(db))
	e.DELETE("/attributes/:id", handlers.DeleteAttribute(db))

	// Routes for question
	e.GET("/question", handlers.GetQuestion(db))

	// Start as a web server
	e.Start(":8000")
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS persons(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		nationality VARCHAR,
		lived VARCHAR,
		city VARCHAR
	);
	CREATE TABLE IF NOT EXISTS attributes(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		person_id INTEGER NOT NULL,
		info TEXT
	);
	`
	_, err := db.Exec(sql)

	if err != nil {
		panic(err)
	}
}
