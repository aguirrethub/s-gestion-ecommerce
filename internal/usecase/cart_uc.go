package usecase

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

/*
CartRepository define el contrato que necesita la capa de casos de uso
para trabajar con carritos.

Principio (arquitectura limpia):
- usecase define interfaces (contratos).
- adapters (memory, db, etc.) implementan esas interfaces.
- así la lógica de aplicación NO depende de infraestructura concreta.
*/
type CartRepository interface {
	// Get devuelve el carrito del cliente. Si no existe, típicamente devuelve uno vacío.
	Get(customerID int) domain.Cart

	// Save persiste el carrito (estado actual) del cliente.
	Save(cart domain.Cart)

	// Clear elimina/vacía el carrito del cliente.
	Clear(customerID int)
}

/*
ProductRepositoryForCart define el contrato mínimo de productos que
necesita el carrito.

¿Por qué existe aparte de ProductRepository?
- Porque el carrito necesita operaciones específicas:
  - buscar producto por ID (para validar existencia/precio/stock)
  - actualizar producto (para descontar stock en un checkout futuro)
*/
type ProductRepositoryForCart interface {
	GetByID(id int) (domain.Product, error)
	Update(p domain.Product) error
}

/*
ViewCart es un caso de uso de consulta.

Responsabilidad:
- Devolver el carrito actual del cliente.
- No aplica reglas complejas: solo delega al repositorio.
*/
func ViewCart(cartRepo CartRepository, customerID int) domain.Cart {
	return cartRepo.Get(customerID)
}

/*
AddProductToCart es un caso de uso de comando (modifica estado).

Responsabilidad:
- Validar que el producto exista.
- Validar cantidad.
- Validar stock suficiente.
- Agregar/actualizar el item en el carrito.
- Persistir el carrito actualizado.

Nota de diseño:
- Aquí se valida stock y cantidad (reglas del negocio en capa aplicación + dominio).
- No se descuenta stock del producto todavía (eso suele hacerse en "checkout").
*/
func AddProductToCart(
	cartRepo CartRepository,
	productRepo ProductRepositoryForCart,
	customerID int,
	productID int,
	quantity int,
) (domain.Cart, error) {

	// 1) Obtener el producto para validar que existe y consultar stock/precio.
	p, err := productRepo.GetByID(productID)
	if err != nil {
		// Si el repo no encuentra el producto, propagamos el error.
		return domain.Cart{}, err
	}

	// 2) Validación de cantidad a nivel de caso de uso (más cerca de la entrada).
	if quantity <= 0 {
		return domain.Cart{}, domain.ErrInvalidQuantity
	}

	// 3) Validación de stock antes de permitir agregar al carrito.
	if p.Stock < quantity {
		return domain.Cart{}, domain.ErrNoStock
	}

	// 4) Obtener el carrito actual del cliente.
	cart := cartRepo.Get(customerID)

	// 5) Aplicar la regla de dominio: agregar item (o acumular si ya existía).
	cart, err = domain.AddItem(cart, domain.CartItem{
		ProductID: p.ID,
		Name:      p.Name,
		Price:     p.Price,
		Quantity:  quantity,
	})
	if err != nil {
		// Por ejemplo: ErrInvalidQuantity (aunque ya validamos antes).
		return domain.Cart{}, err
	}

	// 6) Persistir el carrito actualizado.
	cartRepo.Save(cart)
	return cart, nil
}

/*
RemoveProductFromCart es un caso de uso de comando.

Responsabilidad:
- Obtener el carrito.
- Remover el producto por ID (operación idempotente).
- Guardar el carrito actualizado.

Nota:
- No retorna error porque remover un producto inexistente no es un fallo;
  simplemente no cambia el carrito.
*/
func RemoveProductFromCart(cartRepo CartRepository, customerID int, productID int) domain.Cart {
	cart := cartRepo.Get(customerID)
	cart = domain.RemoveItem(cart, productID)
	cartRepo.Save(cart)
	return cart
}

/*
ClearCart vacía por completo el carrito del cliente.

Responsabilidad:
- Delegar al repositorio la operación de "vaciar".
*/
func ClearCart(cartRepo CartRepository, customerID int) {
	cartRepo.Clear(customerID)
}

/*
CartTotal calcula el total del carrito.

Responsabilidad:
- Obtener el carrito.
- Delegar al dominio el cálculo del total.

Por qué el cálculo está en domain:
- "Total" es una regla del negocio (precio * cantidad).
- Mantenerlo en dominio evita duplicarlo en UI o usecases.
*/
func CartTotal(cartRepo CartRepository, customerID int) float64 {
	cart := cartRepo.Get(customerID)
	return domain.Total(cart)
}
