package view

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Response is the standart format of JSON response
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// HTTPResponse will return HTTP with JSON format
func HTTPResponse(w http.ResponseWriter, statusCode int, statusMessage string, data interface{}) {
	responseJSON, _ := json.Marshal(Response{
		Status:  statusCode,
		Message: statusMessage,
		Data:    data,
	})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(responseJSON)))
	w.WriteHeader(statusCode)

	_, _ = w.Write(responseJSON)
}
