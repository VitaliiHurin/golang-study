// task1 project main.go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type StatusResponse struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func showStatus(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := StatusResponse{Message: "Ok", Timestamp: time.Now().Unix()}
	if json.NewEncoder(w).Encode(response) != nil {
		w.WriteHeader(500)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/status", showStatus)
	log.Fatal(http.ListenAndServe(":8080", router))
}
