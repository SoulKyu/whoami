package handlers

func Login() (c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	authenticated := auth.CheckCredentials(db, username, password)
	session, _ := sessionStore.Get(c.Request(), "session-name")
	session.Values["userID"] = userID
	session.Save(c.Request(), c.Response())

	if authenticated {
		return c.String(http.StatusOK, "Vous êtes connecté !")
	}

	return c.String(http.StatusUnauthorized, "Échec de l'authentification")
}