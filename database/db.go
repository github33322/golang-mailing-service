package database

import (
	"time"

	"github.com/github33322/golang-mailing-service/database/model"
)

type Database interface {
	AutoMigrate() error
	DeleteOldRecords() error
	FindAllCustomers(time time.Time) ([]model.Customer, error)
	CreateCustomer(customer *model.Customer) error
	DeleteCustomer(customer *model.Customer) error
	AllUsersWithMailingID(mailingID int) ([]model.Customer, error)
	DeleteMailingEntry(mailingID int) error
}
