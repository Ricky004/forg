package utils

import (
	"log"
	"os"
)

var (
	logger *log.Logger
)

func init() {
	file, err := os.OpenFile("file-organizer.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file:", err)
	}

	logger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func LogOperation(operation string) {
	logger.Println(operation)
}
