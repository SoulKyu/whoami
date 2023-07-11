package main

import (
	"whoami/routes"
	"whoami/pkg/auth"
	"whoami/pkg/database"

	"github.com/gorilla/sessions"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	var sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))

	// Connexion à la base de données
	db, err := database.GetDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.SessionWithConfig(middleware.SessionConfig{
		SaveToStorage: true,
		Store:         sessionStore,
	}))
	

	e.Static("/", "templates")
	e.Static("/static", "static")

	// Routes
	e.GET("/", routes.Index)
	e.GET("/tutoriels", routes.Tutoriel)
	e.GET("/whoami", routes.Whoami)
	e.POST("/login", )

	e.Logger.Fatal(e.Start(":8080"))
}
