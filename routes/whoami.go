package routes

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func Whoami(c echo.Context) error {
	session, _ := sessionStore.Get(c.Request(), "session-name")

	tmpl := template.Must(template.ParseFiles("templates/whoami.html"))

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		tmpl.Execute(c.Response().Writer, map[string]interface{}{
			"Authenticated": true,
			"Username":      session.Values["username"],
		})
	} else {
		tmpl.Execute(c.Response().Writer, nil)
	}

	return nil
}
