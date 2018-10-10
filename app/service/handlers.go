package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"app/config"
	"app/messaging"

	"github.com/julienschmidt/httprouter"

	"app/models"
	"app/utils"
)

var patientServiceUrl, doctorServiceUrl, appoinmentServiceUrl string

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the main service")
}

// Register a User from the request body
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var user models.User

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	utils.HandleError(err)

	err = r.Body.Close()
	utils.HandleError(err)

	// Save JSON to Post struct
	if err := json.Unmarshal(body, &user); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		err := json.NewEncoder(w).Encode(err)
		utils.HandleError(err)
	}

	// Publish the registration message
	go registeration(user)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func registeration(user models.User) {
	if user.UserType == models.PatientType {
		registerPatient(user)
	} else {
		registerDoctor(user)
	}
}

func registerPatient(user models.User) {
	patient := models.Patient{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, CreatedAt: time.Now().UTC()}
	data, err := json.Marshal(patient)
	utils.HandleError(err)

	err = messaging.Publish(data, "registeration", "patient")
	utils.HandleError(err)
}

func registerDoctor(user models.User) {
	doctor := models.Doctor{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email, Password: user.Password, CreatedAt: time.Now().UTC()}
	data, err := json.Marshal(doctor)
	utils.HandleError(err)

	err = messaging.Publish(data, "registeration", "doctor")
	utils.HandleError(err)
}

// Reserve an appoinment from the request body
func ReserveAppoinment(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var appoinment models.Appoinment

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	utils.HandleError(err)

	err = r.Body.Close()
	utils.HandleError(err)

	// Save JSON to Post struct
	if err := json.Unmarshal(body, &appoinment); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		err := json.NewEncoder(w).Encode(err)
		utils.HandleError(err)
	}

	// Publish the registration message
	go reserveAppoinment(appoinment)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

func reserveAppoinment(appoinment models.Appoinment) {
	appoinment.TimePoint = time.Now()
	data, err := json.Marshal(appoinment)
	utils.HandleError(err)

	err = messaging.Publish(data, "reservation", "appoinment")
	utils.HandleError(err)
}

func FindPatient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	endPointUrl := fmt.Sprintf("%s/patient/%v", patientServiceUrl, id)
	response, err := http.Get(endPointUrl)
	utils.HandleError(err)
	data, errr := ioutil.ReadAll(response.Body)
	utils.HandleError(errr)
	fmt.Fprintf(w, string(data))
}

func FindPatients(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	endPointUrl := fmt.Sprintf("%s/patients", patientServiceUrl)
	response, err := http.Get(endPointUrl)
	utils.HandleError(err)
	data, errr := ioutil.ReadAll(response.Body)
	utils.HandleError(errr)
	fmt.Fprintf(w, string(data))
}

func FindDoctor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	endPointUrl := fmt.Sprintf("%s/doctor/%v", doctorServiceUrl, id)
	response, err := http.Get(endPointUrl)
	utils.HandleError(err)
	data, errr := ioutil.ReadAll(response.Body)
	utils.HandleError(errr)
	fmt.Fprintf(w, string(data))
}

func FindDoctors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	endPointUrl := fmt.Sprintf("%s/doctors", doctorServiceUrl)
	response, err := http.Get(endPointUrl)
	utils.HandleError(err)
	data, errr := ioutil.ReadAll(response.Body)
	utils.HandleError(errr)
	fmt.Fprintf(w, string(data))
}

func FindAllAppoinmentsForDoctor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)
	endPointUrl := fmt.Sprintf("%s/doctor/%v", appoinmentServiceUrl, id)
	fmt.Println(endPointUrl)
	response, err := http.Get(endPointUrl)
	utils.HandleError(err)
	data, errr := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	utils.HandleError(errr)
	fmt.Fprintf(w, string(data))
}

func init() {
	s := config.IniatilizeServicesURLS()
	patientServiceUrl = s.PatientServiceUrl
	doctorServiceUrl = s.DoctorServiceUrl
	appoinmentServiceUrl = s.AppoinmentServiceUrl
}
