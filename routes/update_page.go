package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"whoami/models"
	"whoami/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

func UpdatePageHandler(c echo.Context) error {
	session, _ := sessionStore.Get(c.Request(), "session-name")
	title := c.Param("title")
	fmt.Println("La titre de la page est le suivant : ", title)

	page, err := database.GetPageByTitle(title)
	if err != nil {
		// Si la page n'existe pas, renvoyez une erreur 404.
		return c.String(http.StatusNotFound, "Page not found")
	}
	fmt.Printf("Modification de la page : %s", page.Content)

	// Rendu de la page avec le template HTML et envoie du contenu de la page
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		// Rendu de la page avec le template HTML.
		return c.Render(http.StatusOK, "update.html", map[string]interface{}{
			"Title":         page.Title,
			"Content":       template.HTML(page.Content),
			"Authenticated": true,
			"Username":      session.Values["username"],
		})
	} else {
		// Rendu de la page avec le template HTML.
		return c.Render(http.StatusOK, "update.html", map[string]interface{}{
			"Title":         page.Title,
			"Content":       template.HTML(page.Content),
			"Authenticated": false,
		})
	}

}

func PerformUpdate(c echo.Context) error {
	title := c.FormValue("title")
	unSafeContent := c.FormValue("content")
	p := bluemonday.UGCPolicy()
	content := p.Sanitize(unSafeContent)

	page := &models.Page{
		Title:   title,
		Content: content,
	}

	err := database.UpdatePage(page)
	if err != nil {
		// Si la page n'existe pas, renvoyez une erreur 404.
		return c.String(http.StatusNotFound, "Can't Update page")
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/pages/%s", page.Title))
}
