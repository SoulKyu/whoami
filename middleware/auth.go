// middleware/auth.go
package middleware

import (
	"fmt"
	"net/http"
	"whoami/pkg/session"

	"github.com/labstack/echo/v4"
)

var sessionStore = session.Store

// IsLoggedIn est un middleware qui vérifie si un utilisateur est connecté
func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := sessionStore.Get(c.Request(), "session-name")
		if err != nil {
			fmt.Printf("L'erreur suivante est survenue %v", err)
			return c.String(http.StatusInternalServerError, "Une erreur est survenue lors de la récupération de la session pour IsLoggedIn")
		}

		authenticated, ok := session.Values["authenticated"].(bool)
		if !ok || !authenticated {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}
