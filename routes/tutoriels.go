package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Index route handler
func Tutoriel(c echo.Context) error {
	return c.File("templates/mestutoriels.html")
}
