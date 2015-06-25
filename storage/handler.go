// Handler provides record operating methods for HTTP REST server

package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"

	"golang-study/io"
)

var log = io.GetLogger()

func handleResponse(w http.ResponseWriter, level func(format string, args ...interface{}), code int, message string) {
	w.WriteHeader(code)
	level("Code: %d. Message: %s", code, message)
}

type Handler struct{}

func NewHandler() (h *Handler, err error) {
	h = &Handler{}
	err = io.LoadFromFile(RecordsFilePath, &Records)
	if os.IsNotExist(err) {
		err = nil
		Records = make(RecordsType)
	}
	return
}

func (h *Handler) ShowStatus(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	response := StatusResponse{Message: "Ok", Timestamp: time.Now().Unix()}

	if json.NewEncoder(w).Encode(response) != nil {
		handleResponse(w, log.Error, 500, "Encoding response error")
	} else {
		handleResponse(w, log.Info, 200, "Ok")
	}
}

func (h *Handler) GetRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	record := Record{Key: RecordKeyType(p.ByName("key"))}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if v, err := Records.GetLast(record.Key); err == nil {
		record.Value = v
	} else {
		handleResponse(w, log.Info, 404, "No record found for key '"+string(record.Key)+"'")
		return
	}

	handleResponse(w, log.Info, 200, fmt.Sprintf("%+v", record))

	if json.NewEncoder(w).Encode(record) != nil {
		handleResponse(w, log.Error, 500, "Encoding response error")
	}
}

func (h *Handler) GetRecordHistory(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	recordHistory := RecordHistory{Key: RecordKeyType(p.ByName("key"))}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if v, err := Records.GetAll(recordHistory.Key); err == nil {
		recordHistory.Values = v
	} else {
		handleResponse(w, log.Info, 404, "No record found for key '"+string(recordHistory.Key)+"'")
		return
	}

	handleResponse(w, log.Info, 200, fmt.Sprintf("%+v", recordHistory))

	if json.NewEncoder(w).Encode(recordHistory) != nil {
		handleResponse(w, log.Error, 500, "Encoding response error")
	}
}

func (h *Handler) DeleteRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	key := RecordKeyType(p.ByName("key"))

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if Records.Delete(key) != nil {
		handleResponse(w, log.Info, 404, "No record found for key '"+string(key)+"'")
		return
	}

	if err := io.SaveToFile(RecordsFilePath, &Records); err != nil {
		handleResponse(w, log.Fatalf, 500, "Saving records to file failed")
	}

	handleResponse(w, log.Info, 200, "Record with key '"+string(key)+"' was successfully deleted")
}

func (h *Handler) SetRecord(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	record := Record{Key: RecordKeyType(p.ByName("key"))}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewDecoder(r.Body).Decode(&record)

	if record.Value == "" {
		handleResponse(w, log.Info, 400, "No value for key '"+string(record.Key)+"' received")
		return
	}

	Records.Add(record.Key, record.Value)

	if err := io.SaveToFile(RecordsFilePath, &Records); err != nil {
		handleResponse(w, log.Fatalf, 500, "Saving records to file failed")
	}

	handleResponse(w, log.Info, 201, "Record "+fmt.Sprintf("%+v", record)+" was successfully created")

	if json.NewEncoder(w).Encode(record) != nil {
		handleResponse(w, log.Error, 500, "Encoding response error")
	}
}
