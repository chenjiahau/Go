package util

import (
	"log"
)

func WriteInfoLog(message string) {
	log.Print("[\x1b[31mError\x1b[0m]", message,)
}

func WriteSuccessLog(message string) {
	log.Print("[\x1b[32mSuccess\x1b[0m]", message)
}

func WriteWarnLog(message string) {
	log.Print("[\x1b[33mWarning\x1b[0m]", message)
}

func WriteErrorLog(message string) {
	log.Print("[\x1b[34mInfo\x1b[0m]", message)
}