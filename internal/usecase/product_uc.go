package usecase

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

/*
ProductRepository define el contrato que necesita la capa de casos de uso
para trabajar con productos.

Principio aplicado:
- usecase define interfaces.
- infrastructure (memory, db, etc.) las implementa.
- la lógica de negocio queda desacoplada de la persistencia.
*/
type ProductRepository interface {
	// Create persiste un nuevo producto.
	Create(p domain.Product) error

	// List devuelve todos los productos registrados.
	List() []domain.Product
}

/*
CreateProduct es un caso de uso de comando (modifica estado).

Responsabilidad:
- Validar que el producto cumpla las reglas del dominio.
- Persistir el producto usando el repositorio.

Flujo:
1) Validación del producto (ID, nombre, precio, stock).
2) Persistencia delegada al repositorio.

Nota:
- La CLI no valida productos.
- El repositorio no valida reglas de negocio.
- Este es el punto correcto para coordinar ambas capas.
*/
func CreateProduct(repo ProductRepository, p domain.Product) error {
	// Validación de dominio.
	if err := domain.ValidateProduct(p); err != nil {
		return err
	}

	// Persistencia delegada al repositorio.
	return repo.Create(p)
}

/*
ListProducts es un caso de uso de consulta.

Responsabilidad:
- Obtener todos los productos desde el repositorio.
- No aplica reglas de negocio ni modifica estado.

Diseño:
- Función simple y directa.
- Retorna slice vacío si no hay productos (nunca nil).
*/
func ListProducts(repo ProductRepository) []domain.Product {
	return repo.List()
}
