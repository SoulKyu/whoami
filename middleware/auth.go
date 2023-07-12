// middleware/auth.go
package middleware

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))

// IsLoggedIn est un middleware qui vérifie si un utilisateur est connecté
func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := sessionStore.Get(c.Request(), "session-name")
		userID, ok := session.Values["userID"].(string)
		if !ok || userID == "" {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}