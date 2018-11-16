package utils

import (
	"log"
)

// Logging utils
func HandleError(err error) {
	if err != nil {
		log.Printf(" [ERROR] %s", err)
	}
}
