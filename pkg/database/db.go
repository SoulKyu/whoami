// pkg/database/db.go
package database

import (
	"database/sql"
)

// GetDB retourne une connexion à la base de données SQLite
func GetDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Créez la table des utilisateurs si elle n'existe pas déjà
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			password TEXT
		)
	`)
	if err != nil {
		return nil, err
	}

	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users")
	row.Scan(&count)
	if count == 0 {
		// Si la table est vide, ajoutez l'utilisateur administrateur par défaut
		_, err = db.Exec(`
			INSERT INTO users (username, password) VALUES (?, ?)
		`, "admin", "admin123")
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
