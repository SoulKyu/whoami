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
		dbPath := os.Getenv("DB_PATH")
		if dbPath == "" {
			db, err = sql.Open("sqlite3", "./dev.db")
		} else {
			dbFullPath := fmt.Sprintf("%s/dev.db", dbPath)
			db, err = sql.Open("sqlite3", dbFullPath)
		}

	}

	if err != nil {
		return nil, err
	}

	// Créez la table des utilisateurs si elle n'existe pas déjà
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY,
			username TEXT,
			password TEXT
		);
		CREATE TABLE IF NOT EXISTS pages (
			id INTEGER PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			author TEXT,
			URL TEXT,
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
	stmt, err := db.Prepare("INSERT INTO pages (title, content, url) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécute la requête en passant les données de la page
	_, err = stmt.Exec(page.Title, page.Content, page.URL)
	if err != nil {
		return err
	}

	return nil
}

func UpdatePage(page *models.Page) error {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Prépare une requête SQL pour mettre à jour les données
	stmt, err := db.Prepare("UPDATE pages SET content = ?, URL = ? WHERE title = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Exécute la requête en passant les données de la page
	_, err = stmt.Exec(page.Content, page.URL, page.Title)
	if err != nil {
		return err
	}

	return nil
}

func GetPageById(id int) (*models.Page, error) {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("l'id de la page est le suivant: %d", id)

	// Prépare une requête SQL pour insérer les données
	query := "SELECT content, title, URL FROM pages WHERE id = ?"
	row := db.QueryRow(query, id)

	var content, URL, title string
	err = row.Scan(&content, &title, &URL)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucune ligne trouvée
			log.Printf("no page found with id: %d", id)
			return nil, fmt.Errorf("no page found with id: %d", id)
		}
		// Autre erreur
		log.Printf("an error happens while trying to scan: %v", err)
		return nil, fmt.Errorf("an error happens while trying to scan: %v", err)
	}

	page := &models.Page{
		ID:      id,
		Title:   title,
		Content: content,
		URL:     URL,
	}

	return page, nil

}

func GetPageByTitle(title string) (*models.Page, error) {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Prépare une requête SQL pour insérer les données
	query := "SELECT id, content FROM pages WHERE title = ?"
	row := db.QueryRow(query, title)

	var id int
	var content string
	err = row.Scan(&id, &content)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucune ligne trouvée
			return nil, fmt.Errorf("no page found with title: %s", title)
		}
		// Autre erreur
		return nil, fmt.Errorf("an error happens while trying to scan: %v", err)
	}

	page := &models.Page{
		ID: 	 id,
		Title:   title,
		Content: content,
	}

	return page, nil

}

func DeletePageById(id string) error {
	db, err := GetDB()
	if err != nil {
		log.Fatal(err)
	}

	// Prépare une requête SQL pour supprimer la ligne
	stmt, err := db.Prepare("DELETE FROM pages WHERE id = ?")
	if err != nil {
		log.Fatalf("La page %s n'a pas pu être supprimer : %v", id, err)
	}

	// Execute la requête
	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatalf("La page %s n'a pas pu être supprimer : %v", id, err)
	}

	// Vous pouvez également obtenir le nombre de lignes affectées
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("La page %s n'a pas pu être supprimer car aucune page n'a été trouvée: %v", id, err)
	}

	log.Printf("Deleted %d row(s)", count)

	return nil
}
