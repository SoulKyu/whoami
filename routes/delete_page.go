package routes

import (
	"net/http"
	"whoami/pkg/database"

	"github.com/labstack/echo/v4"
)

func DeletePageByTitle(c echo.Context) error {
	// Obtenir le titre de l'URL
	title := c.Param("title")

	// Supprimer la page de la base de données
	err := database.DeletePageByTitle(title)
	if err != nil {
		// Si une erreur se produit lors de la suppression, renvoyez une erreur 500
		return c.String(http.StatusInternalServerError, "Erreur lors de la suppression de la page")
	}

	// Redirigez l'utilisateur vers la page d'accueil (ou une autre page) après la suppression
	return c.Redirect(http.StatusSeeOther, "/")
}
