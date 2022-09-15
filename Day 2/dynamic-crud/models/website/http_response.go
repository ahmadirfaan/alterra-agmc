package models

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (resp *HTTPResponse) ConvertDataJSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.Code)
	return json.NewEncoder(w).Encode(resp)
}
