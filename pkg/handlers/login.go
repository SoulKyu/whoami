package handlers
import (
	//"database/sql"
	"net/http"
	"whoami/pkg/database"

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

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Vérification des informations d'identification de l'utilisateur...

	authenticated := authenticateUser(username, password)

	if authenticated {
		// Si les informations d'identification sont valides, stockez l'ID de l'utilisateur dans la session
		session, _ := sessionStore.Get(c.Request(), "session-name")
		session.Values["userID"] = username
		session.Save(c.Request(), c.Response())

		return c.Redirect(http.StatusSeeOther, "/")
	}

	return c.String(http.StatusUnauthorized, "Échec de l'authentification")
}

func Logout(c echo.Context) error {
	session, _ := sessionStore.Get(c.Request(), "session-name")
	session.Options.MaxAge = -1 // Supprime la session en définissant MaxAge sur -1
	session.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusSeeOther, "/")
}

// authenticateUser vérifie si l'utilisateur existe dans la base de données
func authenticateUser(username, password string) bool {
	// Connexion à la base de données (à partir de la fonction GetDB dans le package database)
	db, err := database.GetDB()
	if err != nil {
		// Gestion de l'erreur de connexion à la base de données
		return false
	}
	defer db.Close()

	// Exécutez une requête pour vérifier si l'utilisateur existe dans la base de données
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND password = ?", username, password).Scan(&count)
	if err != nil {
		// Gestion de l'erreur lors de l'exécution de la requête
		return false
	}

	// Si l'utilisateur existe dans la base de données, renvoyez true, sinon renvoyez false
	return count > 0
}