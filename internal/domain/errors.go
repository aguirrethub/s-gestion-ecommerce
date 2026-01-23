package domain

import "errors"

/*
Errores de dominio del sistema.

Principio clave:
- Todos los errores definidos aquí representan reglas o estados inválidos
  del NEGOCIO, no errores técnicos.
- Son reutilizados por domain, usecase y adapters.
- Evitan strings “hardcodeados” y comparaciones frágiles en el código.

Este archivo centraliza los errores para mantener coherencia y claridad.
*/
var (
	// =========================
	// ERRORES DE PRODUCTOS
	// =========================

	// ErrInvalidID se usa cuando un ID no es válido:
	// - ID <= 0
	// - ID duplicado
	// - ID inexistente (según el contexto)
	ErrInvalidID = errors.New("ID inválido")

	// ErrEmptyName indica que el nombre del producto está vacío.
	ErrEmptyName = errors.New("nombre vacío")

	// ErrInvalidPrice indica que el precio no cumple reglas básicas del dominio.
	ErrInvalidPrice = errors.New("precio inválido")

	// ErrInvalidStock indica que el stock ingresado es incorrecto
	// (por ejemplo, negativo).
	ErrInvalidStock = errors.New("stock inválido")

	// ErrNoStock indica que no hay stock suficiente para realizar una operación.
	ErrNoStock = errors.New("stock insuficiente")

	// ErrProductNotFound indica que el producto no existe en el sistema.
	// Se usa típicamente al buscar por ID en repositorios.
	ErrProductNotFound = errors.New("producto no encontrado")

	// =========================
	// ERRORES DE CLIENTES
	// =========================

	// ErrInvalidCustomerID indica que el ID del cliente es inválido
	// o que ya existe en el sistema.
	ErrInvalidCustomerID = errors.New("ID de cliente inválido")

	// ErrEmptyCustomerName indica que el nombre del cliente está vacío.
	ErrEmptyCustomerName = errors.New("nombre de cliente vacío")

	// ErrInvalidEmail indica que el email no cumple el formato mínimo válido.
	ErrInvalidEmail = errors.New("email inválido")

	// =========================
	// ERRORES DE CARRITO
	// =========================

	// ErrInvalidQuantity indica que la cantidad ingresada es inválida
	// (menor o igual a cero).
	ErrInvalidQuantity = errors.New("cantidad inválida")

	// ErrEmptyCart indica que se intentó operar sobre un carrito vacío
	// (por ejemplo, checkout sin productos).
	ErrEmptyCart = errors.New("carrito vacío")
)
