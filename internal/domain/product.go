package domain

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

func ValidateProduct(p Product) error {
	if p.ID <= 0 {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	if p.Stock < 0 {
		return ErrInvalidStock
	}
	return nil
}
