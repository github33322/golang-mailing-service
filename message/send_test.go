package message_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/github33322/golang-mailing-service/database"
	"github.com/github33322/golang-mailing-service/email"
	"github.com/github33322/golang-mailing-service/message"
	"github.com/github33322/golang-mailing-service/router"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestSendMessage_Success(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	err := db.InsertTestRecord()
	assert.NoError(t, err)

	body, _ := json.Marshal(map[string]int{"mailing_id": 1})
	req, err := http.NewRequest("POST", "/api/messages/send", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string]string
	err = json.NewDecoder(rr.Body).Decode(&resp)
	assert.NoError(t, err)
	assert.Equal(t, "Successfully processed request.", resp["message"])
}

func TestSendMessage_Fail_NoCustomers(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	// No customers with the given mailingID in the fake DB
	body, _ := json.Marshal(map[string]int{"mailing_id": 999}) // Use an ID that doesn't exist
	req, err := http.NewRequest("POST", "/api/messages/send", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestSendMessage_BadRequest(t *testing.T) {
	db := database.NewFakeDatabase()
	email := email.NewFakeEmail()

	messageService := message.NewMessage(db, email)
	r := mux.NewRouter()

	router.SetUpRoutes(r, messageService)

	// Send an invalid JSON body
	body := []byte(`{bad json}`)
	req, err := http.NewRequest("POST", "/api/messages/send", bytes.NewBuffer(body))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}
