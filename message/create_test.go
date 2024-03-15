package message_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/github33322/golang-mailing-service/database"
	"github.com/github33322/golang-mailing-service/database/model"
	"github.com/github33322/golang-mailing-service/email"
	"github.com/github33322/golang-mailing-service/message"
	"github.com/github33322/golang-mailing-service/router"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateMessage(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	customer := model.Customer{
		Email:     "test@example.com",
		Title:     "Test Title",
		Content:   "Test Content",
		MailingID: 1,
	}
	body, _ := json.Marshal(customer)
	req, err := http.NewRequest("POST", "/api/messages", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseCustomer model.Customer
	err = json.NewDecoder(rr.Body).Decode(&responseCustomer)
	assert.NoError(t, err)
	assert.Equal(t, customer.Email, responseCustomer.Email)
	assert.Equal(t, customer.MailingID, responseCustomer.MailingID)
}
