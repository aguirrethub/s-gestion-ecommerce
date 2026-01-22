package domain

import "errors"

var (
	// Productos
	ErrInvalidID       = errors.New("ID inválido")
	ErrEmptyName       = errors.New("nombre vacío")
	ErrInvalidPrice    = errors.New("precio inválido")
	ErrInvalidStock    = errors.New("stock inválido")
	ErrNoStock         = errors.New("stock insuficiente")
	ErrProductNotFound = errors.New("producto no encontrado")

	// Clientes
	ErrInvalidCustomerID = errors.New("ID de cliente inválido")
	ErrEmptyCustomerName = errors.New("nombre de cliente vacío")
	ErrInvalidEmail      = errors.New("email inválido")

	// Carrito
	ErrInvalidQuantity = errors.New("cantidad inválida")
	ErrEmptyCart       = errors.New("carrito vacío")
)
