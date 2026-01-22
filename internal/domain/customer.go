package domain

type Customer struct {
	ID    int
	Name  string
	Email string
}

func ValidateCustomer(c Customer) error {
	if c.ID <= 0 {
		return ErrInvalidCustomerID
	}
	if c.Name == "" {
		return ErrEmptyCustomerName
	}
	if !isValidEmailBasic(c.Email) {
		return ErrInvalidEmail
	}
	return nil
}

func isValidEmailBasic(email string) bool {
	if email == "" {
		return false
	}

	hasAt := false
	hasDot := false

	for _, ch := range email {
		if ch == '@' {
			hasAt = true
		}
		if ch == '.' {
			hasDot = true
		}
	}

	return hasAt && hasDot
}
