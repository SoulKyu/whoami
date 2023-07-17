package routes

import (
	"html/template"
	"whoami/pkg/session"

	"github.com/labstack/echo/v4"

)

// La clé doit être de 16, 24 ou 32 octets pour sélectionner
// AES-128, AES-192, ou AES-256 respectivement.
var sessionStore = session.Store

func Index(c echo.Context) error {
	// Obtenez la session courante
	session, _ := sessionStore.Get(c.Request(), "session-name")

	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	// Vérifiez si l'utilisateur est authentifié
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
