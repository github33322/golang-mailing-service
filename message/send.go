package message

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/github33322/golang-mailing-service/database/model"
)

func (m *Message) SendMessage(w http.ResponseWriter, r *http.Request) {
	type requestBody struct {
		MailingID int `json:"mailing_id"`
	}

	var body requestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Printf("Error decoding request body: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	err = m.sendEmailAndDelete(body.MailingID)
	if err != nil {
		log.Printf("Error sending emails and deleting records: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Successfully processed request."})
	if err != nil {
		log.Printf("Error encoding response: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}

func (m *Message) sendEmailAndDelete(mailingID int) error {
	customers, err := m.database.AllUsersWithMailingID(mailingID)
	if err != nil {
		log.Printf("Error finding customers: %s\n", err.Error())

		return err
	}

	if len(customers) == 0 {
		err := fmt.Errorf("No customers found with mailing_id: %d", mailingID)
		log.Printf("%s\n", err.Error())

		return err
	}

	for _, customer := range customers {
		err := m.email.SendEmail(customer.Email)
		if err != nil {
			log.Printf("Could not send email: %s\n", err.Error())

			return err
		}

		err = m.deleteRecord(customer)
		if err != nil {
			log.Printf("Could not delete record: %s\n", err.Error())

			return err
		}
	}

	return nil
}

func (m *Message) deleteRecord(customer model.Customer) error {
	err := m.database.DeleteCustomer(&customer)
	if err != nil {
		log.Printf("Could not delete record: %s\n", err.Error())
		return err
	}

	log.Printf("Successfully deleted record for: %s\n", customer.Email)

	return nil
}
