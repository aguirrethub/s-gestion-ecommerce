package usecase

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

type CartRepository interface {
	Get(customerID int) domain.Cart
	Save(cart domain.Cart)
	Clear(customerID int)
}

type ProductRepositoryForCart interface {
	GetByID(id int) (domain.Product, error)
	Update(p domain.Product) error
}

func ViewCart(cartRepo CartRepository, customerID int) domain.Cart {
	return cartRepo.Get(customerID)
}

func AddProductToCart(cartRepo CartRepository, productRepo ProductRepositoryForCart, customerID int, productID int, quantity int) (domain.Cart, error) {
	p, err := productRepo.GetByID(productID)
	if err != nil {
		return domain.Cart{}, err
	}
	if quantity <= 0 {
		return domain.Cart{}, domain.ErrInvalidQuantity
	}
	if p.Stock < quantity {
		return domain.Cart{}, domain.ErrNoStock
	}

	cart := cartRepo.Get(customerID)
	cart, err = domain.AddItem(cart, domain.CartItem{
		ProductID: p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Quantity:  quantity,
	})
	if err != nil {
		return domain.Cart{}, err
	}

	cartRepo.Save(cart)
	return cart, nil
}

func RemoveProductFromCart(cartRepo CartRepository, customerID int, productID int) domain.Cart {
	cart := cartRepo.Get(customerID)
	cart = domain.RemoveItem(cart, productID)
	cartRepo.Save(cart)
	return cart
}

func ClearCart(cartRepo CartRepository, customerID int) {
	cartRepo.Clear(customerID)
}

func CartTotal(cartRepo CartRepository, customerID int) float64 {
	cart := cartRepo.Get(customerID)
	return domain.Total(cart)
}
