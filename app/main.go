package main

import (
	"log"
	"net/http"

	"app/config"
	"app/messaging"
	"app/service"
	"app/utils"
)

const appName = "app-service"

func main() {
	initializeMessaging()

	//  create a new *router instance
	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func initializeMessaging() {
	m := config.IniatilizeMessagingConfig()
	// connecting to rabbitmQ and set it as the MessageStore implementation
	ampqAddress := m.AmpqUrl
	rabbitMq, errr := messaging.NewRabbitMqStore(ampqAddress)
	utils.HandleError(errr)
	messaging.SetMessageStore(rabbitMq)
}
