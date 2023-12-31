package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"log"
	"net/url"
	"whoami/models"
	"whoami/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

func UpdatePageHandler(c echo.Context) error {
	session, _ := sessionStore.Get(c.Request(), "session-name")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		log.Println("failed to convert id to int")
		return err
	}
	log.Printf("Voici la valeur d'ID aprés convertion : %d", id)

	fmt.Println("La titre de la page est le suivant : ", id)

	page, err := database.GetPageById(id)
	if err != nil {
		// Si la page n'existe pas, renvoyez une erreur 404.
		return c.String(http.StatusNotFound, "Page not found")
	}
	fmt.Printf("Modification de la page : %s", page.Title)

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

	currentPage, err := database.GetPageByTitle(title)

	encodedUrl := url.PathEscape(title)
	
	if err != nil {
		// Si la page n'existe pas, renvoyez une erreur 404.
		return c.String(http.StatusNotFound, "Can't Update page")
	}

	page := &models.Page{
		Title:   title,
		Content: content,
		URL: 	 encodedUrl,
	}

	err = database.UpdatePage(page)
	if err != nil {
		// Si la page n'existe pas, renvoyez une erreur 404.
		return c.String(http.StatusNotFound, "Can't Update page")
	}

	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/pages/%d", currentPage.ID))
}
