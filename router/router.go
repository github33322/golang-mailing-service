package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Message interface {
	CreateMessage(w http.ResponseWriter, r *http.Request)
	DeleteMessage(w http.ResponseWriter, r *http.Request)
	SendMessage(w http.ResponseWriter, r *http.Request)
}

func SetUpRoutes(router *mux.Router, message Message) {
	router.HandleFunc("/api/messages", message.CreateMessage).Methods("POST")
	router.HandleFunc("/api/messages/{id}", message.DeleteMessage).Methods("DELETE")
	router.HandleFunc("/api/messages/send", message.SendMessage).Methods("POST")
}
