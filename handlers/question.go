package handlers

import (
	"database/sql"
	"mahi/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetQuestion(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetQuestion(db))
	}
}
