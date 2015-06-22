// HTTP REST server for storing simple key-value records

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"golang-study/io"
	"golang-study/storage"
)

var log = io.GetLogger()

func main() {

	log.Info("Server is starting...")

	router := httprouter.New()
	handler, err := storage.NewHandler()
	if err != nil {
		log.Critical(err.Error())
		return
	}

	router.GET("/status", handler.ShowStatus)
	router.GET("/keys/:key", handler.GetRecord)
	router.GET("/keys/:key/all", handler.GetRecordHistory)
	router.POST("/keys/:key", handler.SetRecord)
	router.DELETE("/keys/:key", handler.DeleteRecord)

	err = http.ListenAndServe(":8080", router)
	log.Critical(err.Error())
}
