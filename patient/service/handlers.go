package service

import (
	"context"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"patient/db"
	"patient/models"
	"patient/utils"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the Patients service")
}

// Register a patient from the request body
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx := r.Context()

	var patient models.Patient

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	utils.HandleError(err)

	if err := r.Body.Close(); err != nil {
		utils.HandleError(err)
	}

	// Save JSON to the patient struct
	if err := json.Unmarshal(body, &patient); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			utils.HandleError(err)
		}
	}

	// hashing password
	patient.Password = getPasswordHash(patient.Password, patient.Email)
	patient.CreatedAt = time.Now()

	// storing at the database
	err = db.RegisterPatient(ctx, patient)
	utils.HandleError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// Retriving the patient with id
func FindPatient(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)

	utils.HandleError(err)
	patient, err := db.FindPatient(ctx, id)
	utils.HandleError(err)

	if err := json.NewEncoder(w).Encode(patient); err != nil {
		utils.HandleError(err)
	}

}

// Retriving the patients list
func FindAllPatients(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	patients, err := db.FindAllPatients(ctx)
	utils.HandleError(err)
	if err := json.NewEncoder(w).Encode(patients); err != nil {
		utils.HandleError(err)
	}

}

func RegisterPatient(s []byte, ctx context.Context) {

	var patient models.Patient
	// Save JSON to the patient struct
	err := json.Unmarshal(s, &patient)
	utils.HandleError(err)

	// hashing password
	patient.Password = getPasswordHash(patient.Password, patient.Email)
	patient.CreatedAt = time.Now()

	// storing at the database
	err = db.RegisterPatient(ctx, patient)
	utils.HandleError(err)
}

func getPasswordHash(password, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}
