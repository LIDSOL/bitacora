package main

import (
	"gitlab.com/lidsol-ng/bitacora/cli"
	"gitlab.com/lidsol-ng/bitacora/database"
)

func main() {
	database.CreateDatabaseFile("bitacora.db")

	cli.PrintMainMenu()
}
