package usecase

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

type ProductRepository interface {
	Create(p domain.Product) error
	List() []domain.Product
}

func CreateProduct(repo ProductRepository, p domain.Product) error {
	if err := domain.ValidateProduct(p); err != nil {
		return err
	}
	return repo.Create(p)
}

func ListProducts(repo ProductRepository) []domain.Product {
	return repo.List()
}
