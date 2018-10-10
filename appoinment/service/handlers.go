package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"appoinment/db"
	"appoinment/models"
	"appoinment/utils"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the appoinments service")
}

// Retriving the appoinments list for a particular doctor id
func FindAllDoctorAppoinments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	id, errr := strconv.Atoi(ps.ByName("id"))
	utils.HandleError(errr)

	appoinments, err := db.FindAllDoctorAppoinments(ctx, id)
	utils.HandleError(err)

	err = json.NewEncoder(w).Encode(appoinments)
	utils.HandleError(err)

}

func RegisterAppoinment(ctx context.Context, s []byte) {

	var appoinment models.Appoinment
	// Save JSON to the appoinment struct
	err := json.Unmarshal(s, &appoinment)
	utils.HandleError(err)
	// storing at the database

	fmt.Println(appoinment)

	err = db.ReserveAppoinment(ctx, appoinment)
	utils.HandleError(err)
}
