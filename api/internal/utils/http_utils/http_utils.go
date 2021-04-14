package http_utils

import (
	"encoding/json"
	"net/http"

	"github.com/cmd-ctrl-q/SELU_ACM/api/internal/utils/errors/rest"
)

func RespondJson(w http.ResponseWriter, method string, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// headers here may not be necessary as cors is implemented in http/rest/handler.go
	if method == "GET" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", method)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err rest.Err) {
	RespondJson(w, "", err.GetStatus(), err)
}
