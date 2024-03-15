package message

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/github33322/golang-mailing-service/database/model"
	"github.com/github33322/golang-mailing-service/validators"
)

func (m *Message) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var customer model.Customer

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		log.Printf("Error decoding request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	ok, err := validators.ValidateCustomer(customer)
	if !ok {
		log.Printf("Invalid customer data: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = m.database.CreateCustomer(&customer)
	if err != nil {
		log.Printf("Error saving customer data to database: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	log.Printf("Successfully created customer with Email: %s, MailingID: %d\n", customer.Email, customer.MailingID)

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		log.Printf("Error encoding response: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
