package message

import (
	"github.com/github33322/golang-mailing-service/database"
	"github.com/github33322/golang-mailing-service/email"
)

type Message struct {
	database database.Database
	email    email.Email
}

func NewMessage(database database.Database, email email.Email) *Message {
	return &Message{
		database: database,
		email:    email,
	}
}
