package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func ResponseSender(w http.ResponseWriter, payload interface{}, status int) {
	response, err := json.Marshal(payload)
	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(status)
	code, err := w.Write(response)
	if err != nil {
		log.Println("Enter: ", code, err)
	}
}
