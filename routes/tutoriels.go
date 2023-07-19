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

	rows, err := db.Query("SELECT id, title, content, URL  FROM pages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var pages []*models.Page

	for rows.Next() {
		page := &models.Page{}
		err := rows.Scan(&page.ID, &page.Title, &page.Content, &page.URL)
		if err != nil {
			log.Fatalf("Erreur logs du scan des pages : %v ", err)
		}
		pages = append(pages, page)
	}

	for _, page := range pages {
		log.Printf("Voici le titre : %s, l'URL : %s et l'id : %d", page.Title, page.URL, page.ID)
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
