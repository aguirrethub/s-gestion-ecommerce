package usecase

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

type CustomerRepository interface {
	Create(c domain.Customer) error
	List() []domain.Customer
}

func CreateCustomer(repo CustomerRepository, c domain.Customer) error {
	if err := domain.ValidateCustomer(c); err != nil {
		return err
	}
	return repo.Create(c)
}

func ListCustomers(repo CustomerRepository) []domain.Customer {
	return repo.List()
}
