package database

import (
	"time"

	"github.com/github33322/golang-mailing-service/database/model"
)

type FakeDB struct {
	Customers []model.Customer
}

func NewFakeDatabase() *FakeDB {
	return &FakeDB{
		Customers: make([]model.Customer, 0),
	}
}

func (f *FakeDB) AutoMigrate() error {
	return nil
}

func (f *FakeDB) DeleteOldRecords() error {
	now := time.Now()
	fiveMinutesAgo := now.Add(-5 * time.Minute)

	var activeCustomers []model.Customer

	for _, customer := range f.Customers {
		if customer.InsertTime.After(fiveMinutesAgo) {
			activeCustomers = append(activeCustomers, customer)
		}
	}

	f.Customers = activeCustomers

	return nil
}

func (f *FakeDB) InsertTestRecord() error {
	testCustomer := model.Customer{
		Email:      "test@example.com",
		Title:      "Test Title",
		Content:    "Test Content",
		MailingID:  1,
		InsertTime: time.Now(),
	}

	f.Customers = append(f.Customers, testCustomer)

	return nil
}

func (f *FakeDB) FindAllCustomers(date time.Time) ([]model.Customer, error) {
	var customers []model.Customer

	for _, customer := range f.Customers {
		if customer.InsertTime.Before(date) {
			customers = append(customers, customer)
		}
	}

	return customers, nil
}

func (f *FakeDB) CreateCustomer(customer *model.Customer) error {
	// Create customer only if pair of email and mailing_id is unique
	for _, existingCustomer := range f.Customers {
		if existingCustomer.Email == customer.Email && existingCustomer.MailingID == customer.MailingID {
			return nil
		}
	}

	f.Customers = append(f.Customers, *customer)

	return nil
}

func (f *FakeDB) AllUsersWithMailingID(mailingID int) ([]model.Customer, error) {
	var customers []model.Customer

	for _, customer := range f.Customers {
		if customer.MailingID == mailingID {
			customers = append(customers, customer)
		}
	}

	return customers, nil
}

func (f *FakeDB) DeleteMailingEntry(mailingID int) error {
	var remainingCustomers []model.Customer

	for _, customer := range f.Customers {
		if customer.MailingID != mailingID {
			remainingCustomers = append(remainingCustomers, customer)
		}
	}

	f.Customers = remainingCustomers

	return nil
}

func (f *FakeDB) DeleteCustomer(target *model.Customer) error {
	var remainingCustomers []model.Customer

	for _, customer := range f.Customers {
		if customer.Email != target.Email {
			remainingCustomers = append(remainingCustomers, customer)
		}
	}

	f.Customers = remainingCustomers

	return nil
}
