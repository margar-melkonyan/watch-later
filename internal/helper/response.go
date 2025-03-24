package helper

import (
	"net/http"
)

type Response struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func SendResponse(
	w http.ResponseWriter,
	err error,
	errCode int,
	code int,
	data []byte,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	if err != nil {
		w.WriteHeader(errCode)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	_, _ = w.Write(data)
}
