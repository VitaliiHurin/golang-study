// task2 project main.go
// HTTP REST server for storing simple key-value records

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"golang-study/task2/controllers"
)

var log = controllers.GetLogger()

func main() {

	log.Info("Server is starting...")

	router := httprouter.New()
	handler, err := controllers.NewHandler()
	if err != nil {
		log.Critical(err.Error())
		return
	}

	router.GET("/status", handler.ShowStatus)
	router.GET("/keys/:key", handler.GetRecord)
	router.POST("/keys/:key", handler.SetRecord)
	router.DELETE("/keys/:key", handler.DeleteRecord)

	err = http.ListenAndServe(":8080", router)
	log.Critical(err.Error())
}
