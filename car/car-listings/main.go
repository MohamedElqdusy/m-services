package main

import (
	"fmt"
	"log"
	"net/http"
	"car-listings/config"
	"car-listings/utils"
	"car-listings/db"
	"car-listings/service"
)

const appName = "car-listings-service"

func main() {
	initiatPostgre()
	//  create a new *router instance
	router := service.NewRouter()
	log.Fatal(http.ListenAndServe(":4422", router))
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

