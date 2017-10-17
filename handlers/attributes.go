package handlers

import (
	"database/sql"
	"mahi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetAttributes(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetAttributes(db))
	}
}

func PutAttribute(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var attribute models.Attribute
		c.Bind(&attribute)
		id, err := models.PutAttribute(db, attribute.Person_id, attribute.Info)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

func DeleteAttribute(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		_, err := models.DeleteAttribute(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}
	}
}
