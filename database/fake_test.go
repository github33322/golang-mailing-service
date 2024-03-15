package database_test

import (
	"testing"
	"time"

	"github.com/github33322/golang-mailing-service/database"
	"github.com/github33322/golang-mailing-service/database/model"
	"github.com/stretchr/testify/assert"
)

func TestFakeDB_AutoMigrate(t *testing.T) {
	db := database.NewFakeDatabase()
	err := db.AutoMigrate()
	assert.NoError(t, err)
}

func TestFakeDB_InsertAndFindCustomers(t *testing.T) {
	db := database.NewFakeDatabase()

	// Insert a test customer
	customer := &model.Customer{
		Email:      "custom@example.com",
		Title:      "Custom Title",
		Content:    "Custom Content",
		MailingID:  2,
		InsertTime: time.Now(),
	}
	err := db.CreateCustomer(customer)
	assert.NoError(t, err)

	// Find all customers before now
	customers, err := db.FindAllCustomers(time.Now().Add(1 * time.Minute))
	assert.NoError(t, err)
	assert.Len(t, customers, 1)

	// Ensure the test record is present
	foundTestRecord := false

	for _, c := range customers {
		if c.Email == customer.Email {
			foundTestRecord = true
			break
		}
	}

	assert.True(t, foundTestRecord, "Test record should be found")
}

func TestFakeDB_DeleteOldRecords(t *testing.T) {
	db := database.NewFakeDatabase()

	// Insert an old customer record
	oldCustomer := &model.Customer{
		Email:      "old@example.com",
		Title:      "Old Title",
		Content:    "Old Content",
		MailingID:  3,
		InsertTime: time.Now(), // 10 minutes ago
	}
	err := db.CreateCustomer(oldCustomer)
	assert.NoError(t, err)

	// Delete old records
	err = db.DeleteOldRecords()
	assert.NoError(t, err)

	// Try to find the old customer
	customers, err := db.FindAllCustomers(time.Now().Add(-1 * time.Minute))
	assert.NoError(t, err)

	// The old customer should not be found
	for _, c := range customers {
		assert.NotEqual(t, "old@example.com", c.Email, "Old customer should have been deleted")
	}
}

func TestFakeDB_DeleteMailingEntry(t *testing.T) {
	db := database.NewFakeDatabase()

	// Insert a couple of customers with the same mailing ID
	customer1 := &model.Customer{
		Email:      "user1@example.com",
		Title:      "User 1",
		Content:    "Content 1",
		MailingID:  4,
		InsertTime: time.Now(),
	}
	customer2 := &model.Customer{
		Email:      "user2@example.com",
		Title:      "User 2",
		Content:    "Content 2",
		MailingID:  4,
		InsertTime: time.Now(),
	}
	err := db.CreateCustomer(customer1)
	assert.NoError(t, err)

	err = db.CreateCustomer(customer2)
	assert.NoError(t, err)

	// Delete mailing entry
	err = db.DeleteMailingEntry(4)
	assert.NoError(t, err)

	// Verify that no customers with the mailing ID exist
	customers, err := db.AllUsersWithMailingID(4)
	assert.NoError(t, err)
	assert.Len(t, customers, 0, "No customers with the mailing ID should exist")
}

func TestFakeDB_DeleteCustomer(t *testing.T) {
	db := database.NewFakeDatabase()

	// Insert a couple of customers
	customer1 := &model.Customer{
		Email:      "user1@example.com",
		Title:      "User 1",
		Content:    "Content 1",
		MailingID:  1,
		InsertTime: time.Now(),
	}
	customer2 := &model.Customer{
		Email:      "user2@example.com",
		Title:      "User 2",
		Content:    "Content 2",
		MailingID:  2,
		InsertTime: time.Now(),
	}
	db.Customers = append(db.Customers, *customer1, *customer2)

	// Delete the first customer
	err := db.DeleteCustomer(customer1)
	assert.NoError(t, err)

	// Verify that only the second customer remains
	assert.Len(t, db.Customers, 1, "There should be exactly one customer remaining in the database")
	assert.Equal(t, "user2@example.com", db.Customers[0].Email, "The remaining customer should be 'user2@example.com'")

	// Optionally, verify that attempting to delete a non-existent customer does not affect the database
	nonExistentCustomer := &model.Customer{Email: "nonexistent@example.com"}
	err = db.DeleteCustomer(nonExistentCustomer)
	assert.NoError(t, err, "Deleting a non-existent customer should not result in an error")
	assert.Len(t, db.Customers, 1, "The number of customers should remain unchanged after attempting to delete a non-existent customer")
}
