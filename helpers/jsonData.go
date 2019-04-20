package helpers

import (
	"net/http"
)

func ResJSON(w http.ResponseWriter, payload []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}
