package handlers

import (
	//"database/sql"
	"net/http"
	"whoami/pkg/auth"
	"whoami/pkg/database"
	"whoami/pkg/session"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
)

var sessionStore = session.Store

func IsLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, err := sessionStore.Get(c.Request(), "session-name")
		if err != nil {
			return err
		}

		auth, ok := session.Values["authenticated"].(bool)
		if !ok || !auth {
			return c.Redirect(http.StatusSeeOther, "/")
		}
		return next(c)
	}
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	db, err := database.GetDB()
	if err != nil {
		return c.String(http.StatusInternalServerError, "Échec de la connexion à la base de données")
	}

	// Vérification des informations d'identification de l'utilisateur...
	authenticated, user, err := auth.CheckCredentials(db, username, password)
	if err != nil {
		return err
	}

	if authenticated {
		// Si les informations d'identification sont valides, stockez l'ID de l'utilisateur dans la session
		session, _ := sessionStore.Get(c.Request(), "session-name")
		session.Values["authenticated"] = true
		session.Values["username"] = user
		err := sessions.Save(c.Request(), c.Response())
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.String(http.StatusUnauthorized, "Échec de l'authentification")
}

func Logout(c echo.Context) error {
	session, err := sessionStore.Get(c.Request(), "session-name")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Une erreur est survenue lors de la récupération de la session pour Logout")
	}

	session.Options.MaxAge = -1 // Supprime la session en définissant MaxAge sur -1
	err = session.Save(c.Request(), c.Response())
	if err != nil {
		return c.String(http.StatusInternalServerError, "Une erreur est survenue lors de la sauvegarde de la session pour Logout")
	}

	return c.Redirect(http.StatusSeeOther, "/")
}
