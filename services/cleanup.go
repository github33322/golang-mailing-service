package services

import (
	"context"
	"log"
	"time"

	"github.com/github33322/golang-mailing-service/database"
)

type CleanupService struct {
	DB database.Database
}

func NewCleanupService(db database.Database) *CleanupService {
	return &CleanupService{
		DB: db,
	}
}

func (c *CleanupService) StartCleanupService(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				c.deleteOldRecords()
			case <-ctx.Done():
				return
			}
		}
	}()
}

func (c *CleanupService) deleteOldRecords() {
	customers, err := c.DB.FindAllCustomers(time.Now().Add(-5 * time.Minute))
	if err != nil {
		log.Println("Could not find customers: ", err.Error())
		return
	}

	if len(customers) == 0 {
		log.Println("No old customers found.")
		return
	}

	// Delete old records
	for _, customer := range customers {
		err := c.DB.DeleteCustomer(&customer)
		if err != nil {
			log.Printf("Error deleting old record for: %s\n", customer.Email)
			continue
		}

		log.Printf("Successfully deleted old record for: %s\n", customer.Email)
	}
}
