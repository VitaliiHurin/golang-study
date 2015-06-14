// task2 project handler.go
// Handler provides record operating methods for HTTP REST server

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"

	"golang-study/task2/models"
)

var log = GetLogger()

func handleResponce(w http.ResponseWriter, level func(format string, args ...interface{}), code int, message string) {
	w.WriteHeader(code)
	level("Code: %d. Message: %s", code, message)
}

type Handler struct{}

func NewHandler() (h *Handler, err error) {
	h = &Handler{}
	err = loadFromFile(models.RecordsFilePath, &models.RecordsMap)
	if os.IsNotExist(err) {
		err = nil
		models.RecordsMap = make(models.RecordsMapType)
	}
	return
}

func (h *Handler) ShowStatus(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := models.StatusResponse{Message: "Ok", Timestamp: time.Now().Unix()}

	if json.NewEncoder(w).Encode(response) != nil {
		handleResponce(w, log.Error, 500, "Encoding response error")
	} else {
		handleResponce(w, log.Info, 200, "Ok")
	}
}

func (h *Handler) GetRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	record := models.Record{Key: models.RecordKeyType(p.ByName("key"))}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if v, ok := models.RecordsMap[record.Key]; ok {
		record.Value = v
	} else {
		handleResponce(w, log.Info, 404, "No record found for key '"+string(record.Key)+"'")
		return
	}

	handleResponce(w, log.Info, 200, fmt.Sprintf("%+v", record))

	if json.NewEncoder(w).Encode(record) != nil {
		handleResponce(w, log.Error, 500, "Encoding response error")
	}
}

func (h *Handler) DeleteRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	key := models.RecordKeyType(p.ByName("key"))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if _, ok := models.RecordsMap[key]; !ok {
		handleResponce(w, log.Info, 404, "No record found for key '"+string(key)+"'")
		return
	}

	delete(models.RecordsMap, key)

	if err := saveToFile(models.RecordsFilePath, &models.RecordsMap); err != nil {
		handleResponce(w, log.Fatalf, 500, "Saving records to file failed")
	}

	handleResponce(w, log.Info, 200, "Record with key '"+string(key)+"' was successfully deleted")
}

func (h *Handler) SetRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	record := models.Record{Key: models.RecordKeyType(p.ByName("key"))}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewDecoder(r.Body).Decode(&record)

	if record.Value == "" {
		handleResponce(w, log.Info, 400, "No value for key '"+string(record.Key)+"' received")
		return
	}

	models.RecordsMap[record.Key] = record.Value

	if err := saveToFile(models.RecordsFilePath, &models.RecordsMap); err != nil {
		handleResponce(w, log.Fatalf, 500, "Saving records to file failed")
	}

	handleResponce(w, log.Info, 201, "Record "+fmt.Sprintf("%+v", record)+" was successfully created")

	if json.NewEncoder(w).Encode(record) != nil {
		handleResponce(w, log.Error, 500, "Encoding response error")
	}
}
