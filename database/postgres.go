package database

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/github33322/golang-mailing-service/database/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PostgresDB struct {
	DBConn *gorm.DB
}

func NewDatabase() Database {
	connectionString := os.Getenv("DATABASE_URL")

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Could not connect to the Postgres database: %s\n", err.Error())
	}

	log.Println("Successfully connected to the Postgres database.")

	return &PostgresDB{
		DBConn: db,
	}
}

func (db *PostgresDB) AutoMigrate() error {
	return db.DBConn.AutoMigrate(&model.Customer{}).Error
}

func (db *PostgresDB) DeleteOldRecords() error {
	now := time.Now()
	fiveMinutesAgo := now.Add(-5 * time.Minute)

	err := db.DBConn.Where("insert_time < ?", fiveMinutesAgo).Delete(&model.Customer{}).Error
	if err != nil {
		log.Printf("Error while deleting old records: %s\n", err)

		return err
	}

	log.Println("Successfully deleted old records.")

	return nil
}

func (db *PostgresDB) FindAllCustomers(date time.Time) ([]model.Customer, error) {
	var customers []model.Customer

	err := db.DBConn.Where("insert_time < ?", date).Find(&customers).Error

	return customers, err
}

func (db *PostgresDB) CreateCustomer(customer *model.Customer) error {
	// Create customer only if pair of email and mailing_id is unique
	var count int

	db.DBConn.Model(&model.Customer{}).Where("email = ? AND mailing_id = ?", customer.Email, customer.MailingID).Count(&count)

	if count > 0 {
		return errors.New("Customer with the same email and mailing_id already exists")
	}

	return db.DBConn.Create(customer).Error
}

func (db *PostgresDB) AllUsersWithMailingID(mailingID int) ([]model.Customer, error) {
	var customers []model.Customer

	err := db.DBConn.Where("mailing_id = ?", mailingID).Find(&customers).Error

	return customers, err
}

func (db *PostgresDB) DeleteMailingEntry(mailingID int) error {
	return db.DBConn.Where("mailing_id = ?", mailingID).Delete(&model.Customer{}).Error
}

func (db *PostgresDB) DeleteCustomer(customer *model.Customer) error {
	return db.DBConn.Delete(customer).Error
}
