package hello

import (
	"log"
	"net/http"
	"encoding/json"
)

func Hello(w http.ResponseWriter, r *http.Request){
	err := json.NewEncoder(w).Encode("Hello from the cloud")
	HandleError(err)
}

func HandleError(err error) {
	if err != nil {
		log.Printf(" [ERROR] %s", err)
	}
}

