package email

type Email interface {
	SendEmail(customerEmail string) error
}
