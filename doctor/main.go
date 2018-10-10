package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"doctor/db"
	"github.com/streadway/amqp"

	"doctor/config"
	"doctor/messaging"
	"doctor/utils"

	"doctor/service"
)

const appName = "doctor-service"

func main() {

	initiatPostgre()
	initiatMessaging()

	//  create a new *router instance
	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":4322", router))
}

func onReceiving(delivery amqp.Delivery) {
	service.RegisterDoctor(delivery.Body, context.Background())
	utils.LogInfo(fmt.Sprintf("Received %v\n", string(delivery.Body)))
}

func initiatPostgre() {
	pc := config.IniatilizePostgreConfig()
	setUpPostgre(pc)
}

func setUpPostgre(pc *config.PostgreConfig) {
	postgersAddress := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",pc.PostgresHost, pc.PostgresPort, pc.PostgresUser, pc.PostgresPassword, pc.PostgresDataBase)
	repository, err := db.NewPostgre(postgersAddress)
	utils.HandleError(err)
	db.SetRepository(repository)
}

func initiatMessaging() {
	mc := config.IniatilizeMessagingConfig()
	setUpRabbitMq(mc)
	setSubscribition()
}

// connecting to rabbitmQ and set it as MessageStore implementation
func setUpRabbitMq(mc *config.MessagingConfig) {
	ampqAddress := mc.AmpqUrl
	rabbitMq, err := messaging.NewRabbitMqStore(ampqAddress)
	utils.HandleError(err)
	messaging.SetMessageStore(rabbitMq)
}

func setSubscribition() {
	err := messaging.Subscribe("registeration", "doctor", "doctor_register_queue", appName, onReceiving)
	utils.FailOnError(err, "Could not start subscribe to register_doctor")
}
