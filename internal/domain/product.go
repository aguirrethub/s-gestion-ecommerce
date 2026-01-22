package domain

import "errors"

type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

var (
	ErrInvalidID    = errors.New("ID inválido")
	ErrEmptyName    = errors.New("nombre vacío")
	ErrInvalidPrice = errors.New("precio inválido")
	ErrInvalidStock = errors.New("stock inválido")
)

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
