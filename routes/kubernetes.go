package routes

import (
	"log"
	"net/http"
	"whoami/pkg/kubernetes"

	"github.com/labstack/echo/v4"
)

func GetKubernetesApplications(c echo.Context) error {
	session, err := sessionStore.Get(c.Request(), "session-name")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Erreur lors de la récupération de la session")
	}

	kubeApplications, err := kubernetes.GetNamespaceResources()
	if err != nil {
		// Si une erreur se produit lors de la suppression, renvoyez une erreur 500
		return c.String(http.StatusInternalServerError, "Erreur lors de la récupération des resources Kubernetes")
	}
	log.Printf("Voici le contenu de kubeApplication : %v", kubeApplications)

	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		// Rendu de la page avec le template HTML.
		return c.Render(http.StatusOK, "kubernetes.html", map[string]interface{}{
			"kubeApplications": kubeApplications,
			"Authenticated":    true,
			"Username":         session.Values["username"],
		})
	} else {
		// Rendu de la page avec le template HTML.
		return c.Redirect(http.StatusUnauthorized, "/")
	}
}
