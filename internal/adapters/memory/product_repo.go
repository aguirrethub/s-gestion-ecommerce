package memory

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

/*
ProductRepo es un repositorio en memoria para productos.

Responsabilidad:
- Almacenar productos mientras el programa está en ejecución.
- Implementar las interfaces de repositorio usadas por los casos de uso:
  - ProductRepository (crear, listar)
  - ProductRepositoryForCart (buscar por ID, actualizar)

Pertenece a la capa de infraestructura (adapters/memory).
*/
type ProductRepo struct {
	// Mapa que asocia un ID de producto con la entidad Product.
	// Key: productID
	// Value: domain.Product
	byID map[int]domain.Product
}

/*
NewProductRepo actúa como constructor del repositorio.

Inicializa el mapa interno que simula una base de datos en memoria.
Se utiliza en main.go para inyectar este repositorio en los casos de uso.
*/
func NewProductRepo() *ProductRepo {
	return &ProductRepo{byID: make(map[int]domain.Product)}
}

/*
Create guarda un nuevo producto en el repositorio.

Reglas aplicadas:
- No se permite crear un producto con un ID ya existente.
- Si el ID está duplicado, se retorna un error de dominio.

Nota:
- Validaciones de negocio más complejas (precio, stock, nombre)
  deben ocurrir en la capa usecase/domain.
*/
func (r *ProductRepo) Create(p domain.Product) error {
	if _, exists := r.byID[p.ID]; exists {
		return domain.ErrInvalidID
	}
	r.byID[p.ID] = p
	return nil
}

/*
List devuelve todos los productos almacenados.

Detalles importantes:
- Retorna un slice, no el map, para no exponer la estructura interna.
- El orden no está garantizado (los maps en Go no mantienen orden).
- Si no hay productos, devuelve un slice vacío.
*/
func (r *ProductRepo) List() []domain.Product {
	out := make([]domain.Product, 0, len(r.byID))
	for _, p := range r.byID {
		out = append(out, p)
	}
	return out
}

/*
GetByID obtiene un producto por su ID.

Este método es fundamental para el carrito:
- Permite verificar que el producto exista.
- Permite consultar precio y stock antes de agregarlo al carrito.

Si el producto no existe, retorna un error de dominio.
*/
func (r *ProductRepo) GetByID(id int) (domain.Product, error) {
	p, ok := r.byID[id]
	if !ok {
		return domain.Product{}, domain.ErrInvalidID
	}
	return p, nil
}

/*
Update actualiza un producto existente.

Uso principal:
- Descontar stock después de una compra o al reservar productos.

Comportamiento:
- Si el producto no existe, retorna error.
- Si existe, reemplaza el valor completo en el repositorio.
*/
func (r *ProductRepo) Update(p domain.Product) error {
	if _, exists := r.byID[p.ID]; !exists {
		return domain.ErrInvalidID
	}
	r.byID[p.ID] = p
	return nil
}
