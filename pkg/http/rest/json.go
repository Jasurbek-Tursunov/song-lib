package rest

import (
	"encoding/json"
	"net/http"
)

func Encode(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
