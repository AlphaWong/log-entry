package utilshttp

import (
	"encoding/json"
	"net/http"
)

func RenderError(w http.ResponseWriter, code int, DetailString string) {
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	errMsg := struct {
		Message string `json:"message"`
		Detail  string `json:"detail,omitempty"`
	}{
		Message: http.StatusText(code),
		Detail:  DetailString,
	}

	json.NewEncoder(w).Encode(errMsg)
}
