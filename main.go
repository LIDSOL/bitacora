package main

import (
	"os"

	"gitlab.com/lidsol-ng/bitacora/cli"
	"gitlab.com/lidsol-ng/bitacora/database"
)

func main() {

	var (
		databaseFile string = "bitacora.db"
		logFile      string = "bitacora.log"
	)

	if os.Getenv("DATABASE_FILE") != "" {
		databaseFile = os.Getenv("DATABASE_FILE")
	}
	if os.Getenv("LOG_FILE") != "" {
		logFile = os.Getenv("LOG_FILE")
	}
	cli.SetLogFile(logFile)

	if os.Getenv("DEBUG") == "true" {
		database.DeleteFile(databaseFile)
		database.DeleteFile(logFile)
		cli.SaveLog("DEBUG mode enabled.")
	}

	cli.SaveLog("Starting program.")
	cli.SaveLog("Database file set to " + databaseFile)
	cli.SaveLog("Log file set to " + logFile)

	cli.LogoLIDSoL()

	if !database.ExistsDatabase(databaseFile) {
		cli.SaveLog("No database file found, creating database.")
		database.CreateDatabaseFile(databaseFile)
	} else {
		cli.SaveLog("Database file found.")
	}

	db, err := database.GetDatabasePointer(databaseFile)
	if err != nil {
		cli.SaveError(err, "")
		return
	}

	var exit bool = false

	for exit == false {
		exit, err = cli.MainMenu(db)
		if err != nil {
			cli.SaveError(err, "")
		}
	}

	defer db.Close()
	cli.SaveLog("Ending program.")
}
