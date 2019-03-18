package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/streadway/amqp"

	config "image-service/config"
	"image-service/db"
	messaging "image-service/messaging"
	"image-service/process"
	utils "image-service/utils"

	service "image-service/service"
)

const appName = "image-service"

func main() {

	initiatRedis()
	initiatMessaging()

	//  create a new *router instance
	router := service.NewRouter()

	log.Fatal(http.ListenAndServe(":4322", router))
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
	err := messaging.Subscribe("images", "resize", "image_resize_queue", appName, onReceiving)
	utils.FailOnError(err, "Could not start subscribe to image_resize")
}

func onReceiving(delivery amqp.Delivery) {
	process.ResizeImageFile(delivery.Body)
}
