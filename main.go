package main

import (
	"context"
	"log"
	"net/http"

	"github.com/github33322/golang-mailing-service/database"
	"github.com/github33322/golang-mailing-service/email"
	"github.com/github33322/golang-mailing-service/message"
	"github.com/github33322/golang-mailing-service/router"
	"github.com/github33322/golang-mailing-service/services"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	db := database.NewDatabase()
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Could not migrate models: %s\n", err.Error())
	}

	email := email.NewFakeEmail()
	message := message.NewMessage(db, email)

	r := mux.NewRouter()

	router.SetUpRoutes(r, message)

	cleanup := services.NewCleanupService(db)
	go cleanup.StartCleanupService(ctx)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

	log.Println("Server is running on port 8080.")
}
