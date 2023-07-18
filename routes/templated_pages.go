package routes

import (
	"html/template"
	"log"
	"net/http"
	"whoami/pkg/database"
	"strconv"

	"github.com/labstack/echo/v4"
)

func TemplatedPages(c echo.Context) error {
	session, _ := sessionStore.Get(c.Request(), "session-name")

	strid := c.Param("id")
	var id *int

	if strid != "" {
		intValue, err := strconv.Atoi(strid)
		if err != nil {
			// Gestion de l'erreur de conversion
			log.Printf("Erreur de conversion de l'ID : %v", err)
		} else {
			id = &intValue
		}
	}
	// Récupérez la page de la base de données.
	// Remplacez ceci par votre propre fonction pour obtenir une page par son titre.
	page, err := database.GetPageById(id)
	if err != nil {
		// Si la page n'existe pas, renvoyez une erreur 404.
		return c.String(http.StatusNotFound, "Page not found")
	}

	// Vérifiez si l'utilisateur est authentifié
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		// Rendu de la page avec le template HTML.
		return c.Render(http.StatusOK, "page.html", map[string]interface{}{
			"Title":         page.Title,
			"URL":			 page.URL,
			"ID": 			 page.ID,
			"Content":       template.HTML(page.Content),
			"Authenticated": true,
			"Username":      session.Values["username"],
		})
	} else {
		// Rendu de la page avec le template HTML.
		return c.Render(http.StatusOK, "page.html", map[string]interface{}{
			"Title":         page.Title,
			"Content":       template.HTML(page.Content),
			"Authenticated": false,
		})
	}
}
