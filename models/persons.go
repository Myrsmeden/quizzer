package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Lived       string `json:"lived"`
	City        string `json:"city"`
}

type PersonCollection struct {
	Persons []Person `json:"items"`
}

func GetPersons(db *sql.DB) PersonCollection {
	sql := "SELECT * FROM persons"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := PersonCollection{}
	for rows.Next() {
		person := Person{}
		err2 := rows.Scan(&person.ID, &person.Name, &person.Nationality, &person.Lived, &person.City)
		if err2 != nil {
			panic(err2)
		}
		result.Persons = append(result.Persons, person)
	}
	return result
}

func PutPerson(db *sql.DB, name string, nationality string, lived string, city string) (int64, error) {
	sql := "INSERT INTO persons(name, nationality, lived, city) VALUES(?,?,?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(name, nationality, lived, city)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeletePerson(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM persons WHERE id = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		panic(err)
	}

	result, err2 := stmt.Exec(id)

	if err2 != nil {
		panic(err2)
	}

	return result.RowsAffected()
}
