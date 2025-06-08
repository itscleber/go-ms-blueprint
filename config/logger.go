package config

import (
	"log"
	"os"
)

var logger *log.Logger

func SetupLogger() {
	logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(message string) {
	logger.SetPrefix("INFO: ")
	logger.Println(message)
}

func Error(message string) {
	logger.SetPrefix("ERROR: ")
	logger.Println(message)
}

func Debug(message string) {
	logger.SetPrefix("DEBUG: ")
	logger.Println(message)
}

func Fatal(message string) {
	logger.SetPrefix("FATAL: ")
	logger.Fatalln(message)
}
