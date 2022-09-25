package ping

import (
	"encoding/json"
	"net/http"
)

// Ping : Return 200
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json; charset=UTF-8")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode("Ping")
}
