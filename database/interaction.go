// database/interaction.go
package database

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

// Know if a user exists
func GetUserExistence(db *sql.DB, userID string) bool {
	query := `
		SELECT userid FROM 'users' WHERE userid = ?;
	`

	if err := db.QueryRow(query, userID).Scan(&userID); err != nil {
		return false
	}

	return true
}

// Add a user to the database
func AddUser(db *sql.DB, u user) error {
	if b := GetUserExistence(db, u.userID); b {
		return errors.New("user already exists")
	}

	query := `
		INSERT INTO 'users' ('userid', 'name', 'surname', 'utype', 'email')
		VALUES (?, ?, ?, ?, ?);
	`

	if _, err := db.Exec(query, u.userID, u.name, u.surname, u.uType, u.email); err != nil {
		return err
	}

	return nil
}

// Get project ID with its name
func GetProjectID(db *sql.DB, name string) (int, error) {
	var id int
	query := `
		SELECT id FROM 'projects' WHERE name = ?;
	`

	if err := db.QueryRow(query, name).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// Add a project to the database
func AddProject(db *sql.DB, p project) error {
	if _, err := GetProjectID(db, p.name); err == nil {
		return errors.New("user already exists")
	}

	query := `
		INSERT INTO 'projects' ('name', 'manager', 'description')
		VALUES (?, ?, ?);
	`

	if _, err := db.Exec(query, p.name, p.manager, p.description); err != nil {
		return err
	}

	return nil
}

// Add a log to the database
func AddLog(db *sql.DB, l log) error {
	query := `
		INSERT INTO 'logs' ('userid', 'projectid')
		VALUES (?, ?);
	`

	if _, err := db.Exec(query, l.userID, l.projectID); err != nil {
		return err
	}

	return nil
}
