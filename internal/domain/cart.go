package domain

type CartItem struct {
	ProductID int
	Name      string
	Price     float64
	Quantity  int
}

type Cart struct {
	CustomerID int
	Items      []CartItem
}

func AddItem(cart Cart, item CartItem) (Cart, error) {
	if item.Quantity <= 0 {
		return cart, ErrInvalidQuantity
	}

	newItems := make([]CartItem, 0, len(cart.Items)+1)
	updated := false

	for _, it := range cart.Items {
		if it.ProductID == item.ProductID {
			it.Quantity += item.Quantity
			updated = true
		}
		newItems = append(newItems, it)
	}

	if !updated {
		newItems = append(newItems, item)
	}

	cart.Items = newItems
	return cart, nil
}

func RemoveItem(cart Cart, productID int) Cart {
	newItems := make([]CartItem, 0, len(cart.Items))
	for _, it := range cart.Items {
		if it.ProductID != productID {
			newItems = append(newItems, it)
		}
	}
	cart.Items = newItems
	return cart
}

func Total(cart Cart) float64 {
	total := 0.0
	for _, it := range cart.Items {
		total += it.Price * float64(it.Quantity)
	}
	return total
}

func IsEmpty(cart Cart) bool {
	return len(cart.Items) == 0
}
