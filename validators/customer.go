package validators

import (
	"fmt"

	"github.com/asaskevich/govalidator"
	"github.com/github33322/golang-mailing-service/database/model"
)

// ValidateCustomer validates the data for a new customer.
func ValidateCustomer(customer model.Customer) (bool, error) {
	if valid := govalidator.IsEmail(customer.Email); !valid {
		return false, fmt.Errorf("Invalid email format")
	}

	if valid := govalidator.IsPrintableASCII(customer.Title); !valid {
		return false, fmt.Errorf("Invalid title format")
	}

	if valid := govalidator.IsPrintableASCII(customer.Content); !valid {
		return false, fmt.Errorf("Invalid content format")
	}

	return true, nil
}
