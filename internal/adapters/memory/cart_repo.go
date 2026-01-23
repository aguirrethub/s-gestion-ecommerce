package memory

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

/*
CartRepo es un repositorio en memoria para carritos.

Responsabilidad:
- Almacenar y recuperar el carrito asociado a cada cliente.
- NO contiene lógica de negocio (no valida stock, no calcula totales).
- Solo persiste el estado del carrito mientras el programa está en ejecución.

Este repositorio implementa la interfaz usecase.CartRepository.
*/
type CartRepo struct {
	// Mapa que asocia un customerID con su carrito.
	// Key: ID del cliente
	// Value: domain.Cart (estado actual del carrito)
	byCustomerID map[int]domain.Cart
}

/*
NewCartRepo actúa como constructor del repositorio.

Inicializa la estructura interna (map) necesaria para
guardar carritos por cliente.

Se usa en main.go para inyectar el repositorio en los casos de uso.
*/
func NewCartRepo() *CartRepo {
	return &CartRepo{byCustomerID: make(map[int]domain.Cart)}
}

/*
Get obtiene el carrito de un cliente específico.

Comportamiento importante:
- Si el cliente NO tiene carrito aún, se devuelve un carrito vacío.
- Nunca devuelve nil, lo que simplifica la lógica de los casos de uso
  y evita validaciones innecesarias en la capa superior.

Este método NO crea efectos secundarios (solo lectura).
*/
func (r *CartRepo) Get(customerID int) domain.Cart {
	cart, ok := r.byCustomerID[customerID]
	if !ok {
		// Si no existe carrito previo, se devuelve uno nuevo y vacío.
		return domain.Cart{
			CustomerID: customerID,
			Items:      []domain.CartItem{},
		}
	}
	return cart
}

/*
Save persiste el carrito de un cliente.

Se utiliza después de:
- agregar productos
- quitar productos
- modificar cantidades

Este método sobreescribe el carrito anterior del cliente.
No valida reglas: asume que el carrito ya fue validado en usecase/domain.
*/
func (r *CartRepo) Save(cart domain.Cart) {
	r.byCustomerID[cart.CustomerID] = cart
}

/*
Clear elimina el contenido del carrito de un cliente.

Implementación:
- No borra la clave del map.
- Reemplaza el carrito por uno vacío.

Esto mantiene un estado consistente y evita tener que
manejar "carrito inexistente" en otros métodos.
*/
func (r *CartRepo) Clear(customerID int) {
	r.byCustomerID[customerID] = domain.Cart{
		CustomerID: customerID,
		Items:      []domain.CartItem{},
	}
}
