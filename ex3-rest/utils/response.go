package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *APIError   `json:"error,omitempty"`
}

type APIError struct {
	Message string `json:"message"`
}

func WriteJSONResponse(writer http.ResponseWriter, statusCode int, data interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	response := JSONResponse{
		Data: data,
	}
	json.NewEncoder(writer).Encode(response)
}

func WriteJSONError(writer http.ResponseWriter, statusCode int, message string) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	response := JSONResponse{
		Error: &APIError{Message: message},
	}
	json.NewEncoder(writer).Encode(response)
}
