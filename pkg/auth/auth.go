// pkg/auth/auth.go
package auth

import (
	"database/sql"
	"whoami/pkg/database"
)

// CheckCredentials vérifie les informations d'identification de l'utilisateur dans la base de données
func CheckCredentials(db *sql.DB, username, password string) bool {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE username = ? AND password = ?", username, password)
	row.Scan(&count)
	return count > 0
}
