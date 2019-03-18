package service

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var imageServiceUrl string

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Hello, welcome to the image processing service")
}
