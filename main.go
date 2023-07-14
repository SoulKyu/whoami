package main

import (
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

func main() {

	//var sessionStore = sessions.NewCookieStore([]byte("your-secret-key"))

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

	e.Static("/", "/app/templates")
	e.Static("/static", "/app/static")

	// Routes
	e.GET("/", routes.Index)
	e.GET("/tutoriels", routes.Tutoriel)
	e.GET("/whoami", routes.Whoami)
	e.POST("/login", handlers.Login)
	e.GET("/logout", handlers.Logout, middleware.IsLoggedIn)

	e.Logger.Fatal(e.Start(":8080"))
}
