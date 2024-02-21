// cli/logs.go
package cli

import (
	"log"
	"os"
)

func getLogFile() string {
	return os.Getenv("LOG_FILE")
}

func SetLogFile(logFile string) {
	os.Setenv("LOG_FILE", logFile)
}

func SaveError(e error, message string) {
	var logFile string = getLogFile()

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Error abriendo archivo de log: %v", err)
	}

	log.SetOutput(f)
	if message == "" {
		log.Println("ERROR: ", e)
	} else {
		log.Println("ERROR: ", e, " ; ", message)
	}

	defer f.Close()
}

func SaveLog(message string) {
	var logFile string = getLogFile()

	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalf("Error abriendo archivo de log: %v", err)
	}

	log.SetOutput(f)
	log.Println("ACTION: ", message)

	defer f.Close()
}
