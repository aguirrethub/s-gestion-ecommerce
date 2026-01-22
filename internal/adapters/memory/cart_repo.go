package memory

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

type CartRepo struct {
	byCustomerID map[int]domain.Cart
}

func NewCartRepo() *CartRepo {
	return &CartRepo{byCustomerID: make(map[int]domain.Cart)}
}

func (r *CartRepo) Get(customerID int) domain.Cart {
	cart, ok := r.byCustomerID[customerID]
	if !ok {
		return domain.Cart{CustomerID: customerID, Items: []domain.CartItem{}}
	}
	return cart
}

func (r *CartRepo) Save(cart domain.Cart) {
	r.byCustomerID[cart.CustomerID] = cart
}

func (r *CartRepo) Clear(customerID int) {
	r.byCustomerID[customerID] = domain.Cart{CustomerID: customerID, Items: []domain.CartItem{}}
}
