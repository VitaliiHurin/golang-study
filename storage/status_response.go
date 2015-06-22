// Provides status response struct

package storage

type StatusResponse struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}
