/**
 * This file handles endpoints for persons
 */
package handlers

import (
	"database/sql"
	"mahi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetPersons(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetPersons(db))
	}
}

func PutPerson(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var person models.Person
		c.Bind(&person)
		id, err := models.PutPerson(db, person.Name, person.Nationality, person.Lived, person.City)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

func DeletePerson(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeletePerson(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
