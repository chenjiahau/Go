package util

import (
	"log"
	"strings"
)

func ParseMessage(message string) []string {
	var messages []string

	if message != "" {
		lines := strings.Split(message, "\n")
		for _, line := range lines {
			if line != "" {
				messages = append(messages, line)
			}
		}
	}

	return messages
}

func WriteInfoLog(message string) {
	log.Print("[\x1b[34mInfo\x1b[0m]", message)
}

func WriteSuccessLog(message string) {
	log.Print("[\x1b[32mSuccess\x1b[0m]", message)
}

func WriteWarnLog(message string) {
	log.Print("[\x1b[33mWarning\x1b[0m]", message)
}

func WriteErrorLog(message string) {
	messages := ParseMessage(message)

	for _, message := range messages {
		log.Print("[\x1b[31mError\x1b[0m]", message)
	}
}