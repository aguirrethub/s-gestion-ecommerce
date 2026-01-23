package usecase

import (
	"time"

	"github.com/aguirrethub/s-gestion-ecommerce/internal/domain"
)

/*
OrderItem representa una línea del detalle del comprobante.
Cada item corresponde a un producto comprado.
*/
type OrderItem struct {
	ProductID int
	Name      string
	UnitPrice float64
	Quantity  int
	LineTotal float64 // UnitPrice * Quantity
}

/*
Order representa el comprobante final de la compra (checkout).
Incluye datos del cliente, detalle de productos y total.
*/
type Order struct {
	ID           string
	CustomerID   int
	CustomerName string
	Items        []OrderItem
	Total        float64
	CreatedAt    time.Time
}

/*
CustomerRepositoryForCheckout permite obtener datos del cliente
sin acoplar el caso de uso a la implementación concreta (memory, DB, etc.).
*/
type CustomerRepositoryForCheckout interface {
	GetByID(id int) (domain.Customer, error)
}

/*
Checkout confirma la compra de un cliente.

Responsabilidades:
1) Obtener el cliente
2) Obtener el carrito
3) Validar carrito no vacío
4) Validar y descontar stock producto por producto
5) Construir el detalle del comprobante
6) Vaciar el carrito
7) Devolver la orden final
*/
func Checkout(
	cartRepo CartRepository,
	productRepo ProductRepositoryForCart,
	customerRepo CustomerRepositoryForCheckout,
	customerID int,
) (Order, error) {

	// Obtener cliente
	customer, err := customerRepo.GetByID(customerID)
	if err != nil {
		return Order{}, err
	}

	// Obtener carrito
	cart := cartRepo.Get(customerID)
	if domain.IsEmpty(cart) {
		return Order{}, domain.ErrEmptyCart
	}

	items := make([]OrderItem, 0, len(cart.Items))
	total := 0.0

	// Procesar cada producto del carrito
	for _, it := range cart.Items {
		p, err := productRepo.GetByID(it.ProductID)
		if err != nil {
			return Order{}, err
		}

		if it.Quantity <= 0 {
			return Order{}, domain.ErrInvalidQuantity
		}

		if p.Stock < it.Quantity {
			return Order{}, domain.ErrNoStock
		}

		// Descontar stock
		p.Stock -= it.Quantity
		if err := productRepo.Update(p); err != nil {
			return Order{}, err
		}

		lineTotal := it.Price * float64(it.Quantity)
		total += lineTotal

		items = append(items, OrderItem{
			ProductID: it.ProductID,
			Name:      it.Name,
			UnitPrice: it.Price,
			Quantity:  it.Quantity,
			LineTotal: lineTotal,
		})
	}

	// Vaciar carrito al completar la compra
	cartRepo.Clear(customerID)

	now := time.Now()

	// Construir orden final
	order := Order{
		ID:           now.Format("20060102150405"),
		CustomerID:   customer.ID,
		CustomerName: customer.Name,
		Items:        items,
		Total:        total,
		CreatedAt:    now,
	}

	return order, nil
}
