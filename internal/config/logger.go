package config

import (
        "log"
        "os"
)

var (

        infoLogger  *log.Logger
        errorLogger *log.Logger
        debugLogger *log.Logger
        fatalLogger *log.Logger
)

func SetupLogger() {
        flags := log.Ldate | log.Ltime | log.Lshortfile
        infoLogger = log.New(os.Stdout, "INFO: ", flags)
        errorLogger = log.New(os.Stdout, "ERROR: ", flags)
        debugLogger = log.New(os.Stdout, "DEBUG: ", flags)
        fatalLogger = log.New(os.Stdout, "FATAL: ", flags)
}

func Info(message string) {
        infoLogger.Println(message)
}

func Error(message string) {
        errorLogger.Println(message)
}

func Debug(message string) {
        debugLogger.Println(message)
}

func Fatal(message string) {
        fatalLogger.Fatalln(message)
}
