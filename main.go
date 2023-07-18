package main

import (
	"html/template"
	"io"
	"whoami/routes"

	//"whoami/pkg/auth"
	"whoami/middleware"
	"whoami/pkg/database"
	"whoami/pkg/handlers"

	//"net/http"

	//"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

type CustomRenderer struct {
	templates *template.Template
}

func (t *CustomRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// Connexion à la base de données
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(echoMiddleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.Static("/", "templates")
	e.Static("/static", "static")
	renderer := &CustomRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Routes
	e.GET("/", routes.Index)
	e.GET("/tutoriels", routes.Tutoriel)
	e.GET("/whoami", routes.Whoami)
	e.POST("/login", handlers.Login)
	e.GET("/logout", handlers.Logout, middleware.IsLoggedIn)
	e.GET("/newPage", routes.CreatePageHandler, middleware.IsLoggedIn)
	e.GET("/kubernetes", routes.GetKubernetesApplications, middleware.IsLoggedIn)
	e.POST("/createPage", routes.CreatePage, middleware.IsLoggedIn)
	e.POST("/deletePage/:title", routes.DeletePageByTitle, middleware.IsLoggedIn)
	e.POST("/updatePage/:title", routes.UpdatePageHandler, middleware.IsLoggedIn)
	e.POST("/kubernetesNamespace/:namespace", routes.UpdatePageHandler, middleware.IsLoggedIn)
	e.POST("/perfromUpdatePage", routes.PerformUpdate, middleware.IsLoggedIn)
	e.GET("/pages/:title", routes.TemplatedPages)
	e.Logger.Fatal(e.Start(":8080"))
}
