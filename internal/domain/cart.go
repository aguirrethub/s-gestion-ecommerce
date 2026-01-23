package domain

/*
CartItem representa un ítem dentro del carrito.

Es una estructura de dominio:
- No sabe nada de bases de datos, CLI o repositorios.
- Solo modela la realidad del negocio: un producto agregado con cantidad.

Cada CartItem es inmutable desde fuera; las modificaciones se hacen
a través de funciones de dominio (AddItem / RemoveItem).
*/
type CartItem struct {
	ProductID int     // Identificador del producto
	Name      string  // Nombre del producto (snapshot al momento de agregar)
	Price     float64 // Precio unitario del producto
	Quantity  int     // Cantidad agregada al carrito
}

/*
Cart representa el carrito de compras de un cliente.

Reglas importantes:
- Un carrito pertenece a un solo cliente (CustomerID).
- Contiene una colección de CartItem.
- El carrito puede existir aunque esté vacío.

Esta entidad vive en el dominio porque modela un concepto central del negocio.
*/
type Cart struct {
	CustomerID int        // Identificador del cliente dueño del carrito
	Items      []CartItem // Ítems actuales del carrito
}

/*
AddItem agrega un producto al carrito.

Comportamiento:
- Si la cantidad es inválida (<= 0), se retorna un error de dominio.
- Si el producto ya existe en el carrito, se incrementa su cantidad.
- Si no existe, se agrega como nuevo ítem.

Diseño:
- No modifica el carrito original directamente.
- Devuelve una nueva versión del carrito (estilo funcional).
- No valida stock ni existencia del producto (eso es responsabilidad del usecase).
*/
func AddItem(cart Cart, item CartItem) (Cart, error) {
	if item.Quantity <= 0 {
		return cart, ErrInvalidQuantity
	}

	// Se crea un nuevo slice para evitar modificar el original directamente.
	newItems := make([]CartItem, 0, len(cart.Items)+1)
	updated := false

	for _, it := range cart.Items {
		if it.ProductID == item.ProductID {
			// Si el producto ya existe, se acumula la cantidad.
			it.Quantity += item.Quantity
			updated = true
		}
		newItems = append(newItems, it)
	}

	// Si el producto no existía en el carrito, se agrega como nuevo ítem.
	if !updated {
		newItems = append(newItems, item)
	}

	cart.Items = newItems
	return cart, nil
}

/*
RemoveItem elimina un producto del carrito por su ProductID.

Comportamiento:
- Si el producto no existe, el carrito queda igual.
- Si existe, se elimina completamente (no reduce cantidad).

Nota:
- No devuelve error porque la operación es idempotente:
  ejecutar varias veces produce el mismo resultado.
*/
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

/*
Total calcula el valor total del carrito.

Regla:
- Suma precio * cantidad de cada ítem.

Este cálculo pertenece al dominio porque define
qué significa "total" en el negocio.
*/
func Total(cart Cart) float64 {
	total := 0.0
	for _, it := range cart.Items {
		total += it.Price * float64(it.Quantity)
	}
	return total
}

/*
IsEmpty indica si el carrito está vacío.

Es una función de conveniencia:
- Simplifica validaciones en capas superiores.
- Evita repetir len(cart.Items) == 0 en múltiples lugares.
*/
func IsEmpty(cart Cart) bool {
	return len(cart.Items) == 0
}
