package service

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"tegaki-service/db"
	"tegaki-service/messaging"
	"tegaki-service/utils"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the main service")
}

// upload an image
func Upload(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	file, reader, err := r.FormFile("img")
	utils.LogInfo(reader.Filename)
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	_, err2 := io.Copy(buf, file)
	utils.HandleError(err2)
	joinChar := []byte("\n")
	s := [][]byte{[]byte(reader.Filename), buf.Bytes()}
	payload := bytes.Join(s, joinChar)
	err = messaging.Publish(payload, "images", "resize")
	utils.HandleError(err)
}

func ShowImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	id := ps.ByName("id")
	// query the database for the image processing state
	imgState, err := db.FindImageRequestStateById(ctx, id)
	utils.HandleError(err)
	if imgState.Processed == true {
		// display the image if processed
		img, err := os.Open("/app/imgs/" + id)
		utils.HandleError(err)
		defer img.Close()
		w.Header().Set("Content-Type", "image/png")
		io.Copy(w, img)
	} else {
		fmt.Fprintf(w, "Not found Object")
	}
}
