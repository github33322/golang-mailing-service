package message

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (m *Message) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	mailingID, err := strconv.Atoi(id)
	if err != nil {
		log.Printf("Error converting id to int: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	users, err := m.database.AllUsersWithMailingID(mailingID)
	if err != nil {
		log.Printf("Error finding customers: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	if len(users) == 0 {
		log.Printf("No customers found with mailing_id: %d", mailingID)
		http.Error(w, "No customers found with mailing_id", http.StatusNotFound)

		return
	}

	err = m.database.DeleteMailingEntry(mailingID)
	if err != nil {
		log.Printf("Error deleting customer: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	log.Printf("Successfully deleted customer with mailing_id: %d\n", mailingID)
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(map[string]string{"message": "Successfully deleted customer."})
	if err != nil {
		log.Printf("Error encoding response: %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
