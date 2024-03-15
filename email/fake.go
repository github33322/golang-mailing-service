package email

import "log"

type FakeEmail struct {
}

func NewFakeEmail() Email {
	return &FakeEmail{}
}

func (e *FakeEmail) SendEmail(customerEmail string) error {
	log.Printf("Successfully sent email to: %s\n", customerEmail)
	return nil
}
