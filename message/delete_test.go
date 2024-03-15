package message_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/github33322/golang-mailing-service/database"
	"github.com/github33322/golang-mailing-service/database/model"
	"github.com/github33322/golang-mailing-service/email"
	"github.com/github33322/golang-mailing-service/message"
	"github.com/github33322/golang-mailing-service/router"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestDeleteMessage_Success(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	// Pre-populate the fake database with a customer that has a specific mailingID
	db.Customers = append(db.Customers, model.Customer{
		Email:      "test@example.com",
		Title:      "Test",
		Content:    "Content",
		MailingID:  1,
		InsertTime: time.Now(),
	})

	req, _ := http.NewRequest("DELETE", "/api/messages/1", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err := json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Successfully deleted customer.", response["message"])

	// Verify the customer was deleted
	customers, err := db.AllUsersWithMailingID(1)
	assert.NoError(t, err)
	assert.Empty(t, customers)
}

func TestDeleteMessage_NotFound(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	req, _ := http.NewRequest("DELETE", "/api/messages/99", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestDeleteMessage_BadRequest(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	req, _ := http.NewRequest("DELETE", "/api/messages/notanumber", nil)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
