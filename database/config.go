// database/config.go
package database

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// --- Structures for the database

type User struct {
	// userID: Account number, RFC
	userID  string
	name    string
	surname string
	// utype: student, worker
	uType string
	email string
}

type Project struct {
	name        string
	manager     string
	description string
}

type Log struct {
	userID    string
	projectID int
}

var (
	userTypes = []string{"student", "worker"}
)

// Basic database operations

func GetDatabasePointer(filepath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ExistsDatabase(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateDatabaseFile(filePath string) error {
	var err error
	db, err := sql.Open("sqlite3", filePath)

	if err != nil {
		return err
	}

	query := `
		CREATE TABLE IF NOT EXISTS 'users' (
			'id' INTEGER PRIMARY KEY,
			'userid' VARCHAR(13) NOT NULL,
			'name' VARCHAR(128) NOT NULL,
			'surname' VARCHAR(128) NOT NULL,
			'utype' VARCHAR(32) NOT NULL,
			'email' VARCHAR(256) NOT NULL
		);
	`
	if _, err = db.Exec(query); err != nil {
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS 'projects' (
			'id' INTEGER PRIMARY KEY,
			'name' VARCHAR(128) NOT NULL,
			'manager' VARCHAR(13) NOT NULL,
			'description' VARCHAR(1024) NOT NULL,
			FOREIGN KEY ('manager') REFERENCES 'users'('userid')
		);
	`
	if _, err = db.Exec(query); err != nil {
		return err
	}

	query = `
		CREATE TABLE IF NOT EXISTS 'logs' (
			'num' INTEGER PRIMARY KEY,
			'userid' VARCHAR(13) NOT NULL,
			'projectid' INT NOT NULL,
			'inTime' DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY ('userid') REFERENCES 'users'('userid'),
			FOREIGN KEY ('projectid') REFERENCES 'projects'('id')
		);
	`
	if _, err = db.Exec(query); err != nil {
		return err
	}

	if err = db.Close(); err != nil {
		return err
	}

	return nil
}

func DeleteFile(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}
