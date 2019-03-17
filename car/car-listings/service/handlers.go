package service

import (

	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"car-listings/utils"
	"car-listings/db"
	"car-listings/models"
	"car-listings/parsing"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the car-listings service")
}

// save cars from the request body
func SaveCars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx := r.Context()

	var cars []models.Car

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	utils.HandleError(err)

	if err := r.Body.Close(); err != nil {
		utils.HandleError(err)
	}

	// Save JSON to the cars struct
	if err := json.Unmarshal(body, &cars); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		err = json.NewEncoder(w).Encode(err)
		utils.FailOnError(err, "with body" + string(body))
		return 
	}

	// storing at the database
	for _, car := range cars {
		// 0 is dealer_id for the rest of the providers
		err = db.StoreCar(ctx, car, 0)
		utils.HandleError(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// save car given dealer_id from CSV file
func saveCarByDealer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	file, reader, err := r.FormFile("csv")
	utils.LogInfo(reader.Filename)
	defer file.Close() 
	ctx := r.Context()
	dealerId, err := strconv.ParseUint(ps.ByName("dealer_id"), 10, 64)
	utils.HandleError(err)
	
	cars, err := parsing.ReadCSVFromHttpRequest(file)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		err = json.NewEncoder(w).Encode(err)
		utils.HandleError(err)
		return
	}
	// storing at the database
	for _, car := range cars {
		err = db.StoreCar(ctx, car, dealerId)
		utils.HandleError(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func SearchCarsByMake(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	make := ps.ByName("make")
	cars, err := db.FindCarsByMake(ctx, make)
	utils.HandleError(err)
	err = json.NewEncoder(w).Encode(cars)
	utils.HandleError(err)
}

func SearchCarsByModel(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	model := ps.ByName("model")
	cars, err := db.FindCarsByModel(ctx, model)
	utils.HandleError(err)
	err = json.NewEncoder(w).Encode(cars)
	utils.HandleError(err)
}

func SearchCarsByYear(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	year, err := strconv.ParseUint(ps.ByName("year"), 10, 64)
	cars, err := db.FindCarsByYear(ctx, year)
	utils.HandleError(err)
	err = json.NewEncoder(w).Encode(cars)
	utils.HandleError(err)
}

func SearchCarsByColor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	color := ps.ByName("color")
	cars, err := db.FindCarsByColor(ctx, color)
	utils.HandleError(err)
	err = json.NewEncoder(w).Encode(cars)
	utils.HandleError(err)
}

func AllCars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	cars, err := db.FindAllCars(ctx)
	utils.HandleError(err)
	err = json.NewEncoder(w).Encode(cars)
	utils.HandleError(err)
}