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

	"doctor/utils"
	"doctor/db"
	"doctor/models"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the doctors service")
}

// Register a doctor from the request body
func Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	ctx := r.Context()

	var doctor models.Doctor

	// Read request body and close it
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	utils.HandleError(err)

	if err := r.Body.Close(); err != nil {
		utils.HandleError(err)
	}

	// Save JSON to the doctor struct
	if err := json.Unmarshal(body, &doctor); err != nil {

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)

		err = json.NewEncoder(w).Encode(err)
		utils.HandleError(err)

	}

	// hashing password
	doctor.Password = getPasswordHash(doctor.Password, doctor.Email)
	doctor.CreatedAt = time.Now()

	// storing at the database
	err = db.RegisterDoctor(ctx, doctor)
	utils.HandleError(err)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
}

// Retriving the doctor with id
func FindDoctor(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	id, err := strconv.ParseUint(ps.ByName("id"), 10, 64)

	utils.HandleError(err)
	doctor, err := db.FindDoctor(ctx, id)
	utils.HandleError(err)

	err = json.NewEncoder(w).Encode(doctor)
	utils.HandleError(err)

}

// Retriving the doctor list
func FindAllDoctors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	doctors, err := db.FindAllDoctors(ctx)
	utils.HandleError(err)
	err = json.NewEncoder(w).Encode(doctors)
	utils.HandleError(err)
}

func RegisterDoctor(s []byte, ctx context.Context) {

	var doctor models.Doctor
	// Save JSON to the doctor struct
	err := json.Unmarshal(s, &doctor)
	utils.HandleError(err)

	// hashing password
	doctor.Password = getPasswordHash(doctor.Password, doctor.Email)
	doctor.CreatedAt = time.Now()

	// storing at the database
	err = db.RegisterDoctor(ctx, doctor)
	utils.HandleError(err)
}

func getPasswordHash(password, salt string) string {
	h := sha1.New()
	io.WriteString(h, salt)
	io.WriteString(h, password)
	return fmt.Sprintf("%x", h.Sum(nil))
}
