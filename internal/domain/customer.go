package domain

import "errors"

type Customer struct {
	ID    int
	Name  string
	Email string
}

var (
	ErrInvalidCustomerID = errors.New("ID de cliente inválido")
	ErrEmptyCustomerName = errors.New("nombre de cliente vacío")
	ErrInvalidEmail      = errors.New("email inválido")
)

func ValidateCustomer(c Customer) error {
	if c.ID <= 0 {
		return ErrInvalidCustomerID
	}
	if c.Name == "" {
		return ErrEmptyCustomerName
	}
	if c.Email == "" || !containsAt(c.Email) {
		return ErrInvalidEmail
	}
	return nil
}

func containsAt(email string) bool {
	for _, ch := range email {
		if ch == '@' {
			return true
		}
	}
	return false
}
