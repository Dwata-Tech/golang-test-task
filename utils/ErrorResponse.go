package utils

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"message"`
}

// RespondJSON makes the response with payload as json format
func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			//logger.Fatal(err.Error())
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(response)
	if err != nil {
		//logger.Fatal(err.Error())
	}
}

// RespondError makes the error response with payload as json format
func RespondError(w http.ResponseWriter, status int, err string) {
	//log.Errorln(err, "session ID:", sessionID)
	RespondJSON(w, status, ErrorResponse{Error: err})
}
