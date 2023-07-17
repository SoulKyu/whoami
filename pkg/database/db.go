// pkg/database/db.go
package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	models "whoami/models"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// GetDB retourne une connexion à la base de données SQLite
func GetDB() (*sql.DB, error) {
	var db *sql.DB
	var err error

	if os.Getenv("APP_ENV") == "production" {
		db, err = sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	} else {
		db, err = sql.Open("sqlite3", "./dev.db")
	}

	if err != nil {
		return nil, err
	}

	// Créez la table des utilisateurs si elle n'existe pas déjà
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username TEXT,
			password TEXT
		);
		CREATE TABLE IF NOT EXISTS pages (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			author TEXT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
		adminPassword := os.Getenv("ADMIN_PASSWORD")
		if adminPassword == "" {
			adminPassword = "admin123"
		}
		_, err = db.Exec(`
			INSERT INTO users (username, password) VALUES (?, ?)
		`, "admin", adminPassword)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}

func CreatePage(page *models.Page) error {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Prépare une requête SQL pour insérer les données
	stmt, err := db.Prepare("INSERT INTO pages (title, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécute la requête en passant les données de la page
	_, err = stmt.Exec(page.Title, page.Content)
	if err != nil {
		return err
	}

	return nil
}

func GetPageByTitle(title string) (*models.Page, error) {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Prépare une requête SQL pour insérer les données
	query := "SELECT content FROM pages WHERE title = ?"
	row := db.QueryRow(query, title)

	var content string
	err = row.Scan(&content)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucune ligne trouvée
			return nil, fmt.Errorf("no page found with title: %s", title)
		}
		// Autre erreur
		return nil, err
	}

	page := &models.Page{
		Title:   title,
		Content: content,
	}

	return page, nil

}

func DeletePageByTitle(title string) error {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Prépare une requête SQL pour supprimer la ligne
	stmt, err := db.Prepare("DELETE FROM pages WHERE title = ?")
	if err != nil {
		log.Fatalf("La page %s n'a pas pu être supprimer : %v", title, err)
	}

	// Execute la requête
	res, err := stmt.Exec(title)
	if err != nil {
		log.Fatalf("La page %s n'a pas pu être supprimer : %v", title, err)
	}

	// Vous pouvez également obtenir le nombre de lignes affectées
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("La page %s n'a pas pu être supprimer car aucune page n'a été trouvée: %v", title, err)
	}

	log.Printf("Deleted %d row(s)", count)

	return nil
}
