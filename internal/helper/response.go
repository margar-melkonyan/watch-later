package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data     any               `json:"data"`
	Messages map[string]string `json:"messages"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func SendResponse(
	w http.ResponseWriter,
	code int,
	data Response,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}

func SendError(w http.ResponseWriter, code int, message MessageResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response, err := json.Marshal(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(response)
}
