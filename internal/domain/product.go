package domain

/*
Product representa un producto disponible para la venta.

Es una entidad de dominio:
- Modela un concepto central del negocio.
- No conoce nada de cómo se guarda (repositorio) ni cómo se muestra (CLI).
- Contiene solo los datos esenciales del producto.
*/
type Product struct {
	ID    int     // Identificador único del producto
	Name  string  // Nombre del producto
	Price float64 // Precio unitario del producto
	Stock int     // Cantidad disponible en inventario
}

/*
ValidateProduct valida las reglas básicas del dominio para un producto.

Responsabilidad:
- Garantizar que un Product tenga valores coherentes antes de ser creado
  o utilizado en los casos de uso.

Reglas de dominio:
- El ID debe ser mayor que 0.
- El nombre no puede estar vacío.
- El precio debe ser mayor que 0.
- El stock no puede ser negativo.

Nota:
- Esta función NO persiste el producto.
- Solo valida las reglas del negocio.
*/
func ValidateProduct(p Product) error {
	if p.ID <= 0 {
		return ErrInvalidID
	}
	if p.Name == "" {
		return ErrEmptyName
	}
	if p.Price <= 0 {
		return ErrInvalidPrice
	}
	if p.Stock < 0 {
		return ErrInvalidStock
	}
	return nil
}
