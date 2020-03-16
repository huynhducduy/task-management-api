package app

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type MessageResponse struct {
	Message string `json:"message"`
}

func responseInternalError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(MessageResponse{
		Message: "Internal error!",
	})
	log.Printf(err.Error() + "\n" + string(debug.Stack()))
}

func responseCustomError(w http.ResponseWriter, httpCode int, message string) {
	w.WriteHeader(httpCode)
	json.NewEncoder(w).Encode(MessageResponse{
		Message: message,
	})
}

func responseOK(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func logg(x interface{}) {
	fmt.Printf("%+v\n", x)
}
