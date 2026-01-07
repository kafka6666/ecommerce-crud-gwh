package utils

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data any, statusCode int) {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(b); err != nil {
		// write failure after headers sent — can't send another header here
		return
	}
}
