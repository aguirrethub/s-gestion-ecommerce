package memory

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

type ProductRepo struct {
	byID map[int]domain.Product
}

func NewProductRepo() *ProductRepo {
	return &ProductRepo{byID: make(map[int]domain.Product)}
}

func (r *ProductRepo) Create(p domain.Product) error {
	if _, exists := r.byID[p.ID]; exists {
		return domain.ErrInvalidID
	}
	r.byID[p.ID] = p
	return nil
}

func (r *ProductRepo) List() []domain.Product {
	out := make([]domain.Product, 0, len(r.byID))
	for _, p := range r.byID {
		out = append(out, p)
	}
	return out
}

// 👇 ESTOS DOS MÉTODOS SON NUEVOS

func (r *ProductRepo) GetByID(id int) (domain.Product, error) {
	p, ok := r.byID[id]
	if !ok {
		return domain.Product{}, domain.ErrInvalidID
	}
	return p, nil
}

func (r *ProductRepo) Update(p domain.Product) error {
	if _, exists := r.byID[p.ID]; !exists {
		return domain.ErrInvalidID
	}
	r.byID[p.ID] = p
	return nil
}
