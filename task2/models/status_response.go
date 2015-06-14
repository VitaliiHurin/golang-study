// task2 project status_response.go
// Provides status response struct

package models

type StatusResponse struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
