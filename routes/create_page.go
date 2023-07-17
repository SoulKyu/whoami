package routes

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"whoami/models"
	"whoami/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/microcosm-cc/bluemonday"
)

func CreatePageHandler(c echo.Context) error {
	// Obtenez la session courante
	session, _ := sessionStore.Get(c.Request(), "session-name")

	tmpl := template.Must(template.ParseFiles("templates/create.html"))

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

func CreatePage(c echo.Context) error {
	// parse form values
	title := c.FormValue("title")
	unSafeContent := c.FormValue("content")
	p := bluemonday.UGCPolicy()
	content := p.Sanitize(unSafeContent)

	// create a new page
	page := &models.Page{
		Title:   title,
		Content: content,
	}

	// save page to database
	err := database.CreatePage(page)
	if err != nil {
		log.Fatal(err)
	}

	// redirect to new page
	return c.Redirect(http.StatusSeeOther, fmt.Sprintf("/pages/%d", page.ID))
}
