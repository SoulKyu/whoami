package routes

import (
	"github.com/labstack/echo/v4"
)

// Index route handler
func Index(c echo.Context) error {
	return c.File("templates/index.html")
}
