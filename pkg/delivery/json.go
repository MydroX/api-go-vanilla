package delivery

import (
	"encoding/json"
	"net/http"
)

type httpResponse struct {
	Code    int
	Message string
}

type httpResponseBody struct {
	Code    int
	Message string
	Body    interface{}
}

func JSONError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	resp := httpResponse{
		Code:    code,
		Message: message,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func JSONResponse(w http.ResponseWriter, arg interface{}, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	resp := httpResponse{
		Code:    code,
		Message: message,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
