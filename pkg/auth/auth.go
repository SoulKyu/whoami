package auth

import (
	"database/sql"
	"fmt"
)

// CheckCredentials vérifie les informations d'identification de l'utilisateur dans la base de données
func CheckCredentials(db *sql.DB, username, password string) (bool, string, error) {
	var count int
	var user string
	row := db.QueryRow("SELECT COUNT(*), username FROM users WHERE username = ? AND password = ?", username, password)
	err := row.Scan(&count, &user)

	if err != nil {
		return false, "", err
	}

	if count > 0 {
		fmt.Printf("L'utilisateur suivant : %s a été identifié \n", user)
		return true, user, nil
	}

	return false, "", nil
}
