package memory

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

type CustomerRepo struct {
	byID map[int]domain.Customer
}

func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{byID: make(map[int]domain.Customer)}
}

func (r *CustomerRepo) Create(c domain.Customer) error {
	if _, exists := r.byID[c.ID]; exists {
		return domain.ErrInvalidCustomerID // simple: ID repetido
	}
	r.byID[c.ID] = c
	return nil
}

func (r *CustomerRepo) List() []domain.Customer {
	out := make([]domain.Customer, 0, len(r.byID))
	for _, c := range r.byID {
		out = append(out, c)
	}
	return out
}
