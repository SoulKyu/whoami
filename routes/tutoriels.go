package routes

import (
	"html/template"
	"log"
	"whoami/models"
	"whoami/pkg/database"

	"github.com/labstack/echo/v4"
)

func Tutoriel(c echo.Context) error {
	session, _ := sessionStore.Get(c.Request(), "session-name")
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT title, content FROM pages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pages []*models.Page

	for rows.Next() {
		page := &models.Page{}
		err := rows.Scan(&page.Title, &page.Content)
		if err != nil {
			log.Fatal(err)
		}
		pages = append(pages, page)
	}

	tmpl := template.Must(template.ParseFiles("templates/mestutoriels.html"))

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		tmpl.Execute(c.Response().Writer, map[string]interface{}{
			"Authenticated": true,
			"Username":      session.Values["username"],
			"Pages":         pages,
		})
	} else {
		tmpl.Execute(c.Response().Writer, map[string]interface{}{
			"Pages": pages,
		})
	}

	return nil
}
