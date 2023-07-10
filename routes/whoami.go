package routes

import (
	"github.com/labstack/echo/v4"
)

// Index route handler
func Whoami(c echo.Context) error {
	return c.File("templates/whoami.html")
}
