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

// JSONResponse is a helper function that returns a JSON response
func JSONResponse(w http.ResponseWriter, code int, message string) {
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

// JSONResponseWithBody is a helper function that returns a JSON response with a body
func JSONResponseWithBody(w http.ResponseWriter, arg interface{}, code int, message string) {
	w.Header().Set("Content-Type", "application/json")

	resp := httpResponseBody{
		Code:    code,
		Message: message,
		Body:    arg,
	}

	res, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
}
