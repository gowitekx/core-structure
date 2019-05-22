package v1

import (
	"encoding/json"
	"io"
	"net/http"
)

//jsonResponse Type
type jsonResponse struct {
	// Reserved field to add some meta information to the API response
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

//jsonErrorResponse Type
type jsonErrorResponse struct {
	Status  int    `json:"status"`
	Success bool   `json:"success"`
	Error   string `json:"message"`
}

//Reads Input from the body
func ReadInput(rBody io.ReadCloser, input interface{}) error {
	decoder := json.NewDecoder(rBody)
	err := decoder.Decode(input)
	return err
}

// Writes the response as a standard JSON response with StatusOK
func WriteOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&jsonResponse{Success: true, Status: http.StatusOK, Data: m}); err != nil {
		WriteErrorResponse(w, http.StatusInternalServerError, "Internal Sever Error!")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).Encode(&jsonErrorResponse{Success: false, Status: errorCode, Error: errorMsg})
}
