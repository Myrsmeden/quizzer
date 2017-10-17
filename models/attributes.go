package models

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Attribute struct {
	ID        int    `json:"id"`
	Person_id int    `json:"person_id"`
	Info      string `json:"info"`
	Person    string `json:"person"`
}

type AttributeCollection struct {
	Attributes []Attribute `json:"items"`
}

func GetAttributes(db *sql.DB) AttributeCollection {
	sql := "SELECT attributes.ID, attributes.person_id, attributes.info, persons.name as person FROM attributes JOIN persons ON attributes.person_id = persons.ID"
	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	result := AttributeCollection{}
	for rows.Next() {
		attribute := Attribute{}
		err2 := rows.Scan(&attribute.ID, &attribute.Person_id, &attribute.Info, &attribute.Person)
		if err2 != nil {
			panic(err2)
		}
		result.Attributes = append(result.Attributes, attribute)
	}
	return result
}

func PutAttribute(db *sql.DB, person_id int, info string) (int64, error) {
	sql := "INSERT INTO attributes(person_id, info) VALUES(?,?)"

	stmt, err := db.Prepare(sql)
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	result, err2 := stmt.Exec(person_id, info)

	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

func DeleteAttribute(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM attributes WHERE id = ?"
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
