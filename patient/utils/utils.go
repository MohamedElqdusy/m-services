package utils

import (
	"fmt"
	"log"
)

// Logging utils
func HandleError(err error) {
	if err != nil {
		log.Printf(" [ERROR] %s", err)
	}
}

func LogInfo(s string) {
	log.Printf(" [INFO] %s", s)
}

func FailOnError(err error, msg string) {
	if err != nil {
		fmt.Printf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
