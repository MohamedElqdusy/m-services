package main

import (
	"log"
	"net/http"
	"strconv"

	"tegaki-service/config"
	"tegaki-service/db"
	"tegaki-service/messaging"
	"tegaki-service/service"
	"tegaki-service/utils"
)

const appName = "tegaki-service"

func main() {
	initiatRedis()
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

func initiatRedis() {
	r := config.IniatilizeRedisConfig()
	setUpRedis(r)
}

func setUpRedis(rc *config.RedisConfig) {
	url := rc.RedisAdress + ":" + rc.RedisPort
	redisDatabase, err := strconv.Atoi(rc.RedisDataBase)
	redis, err := db.NewRedis(url, rc.RedisPassword, redisDatabase)
	utils.HandleError(err)
	db.SetRepository(*redis)
}
