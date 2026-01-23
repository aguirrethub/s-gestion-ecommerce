package memory

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

/*
CustomerRepo es un repositorio en memoria para clientes.

Responsabilidad:
- Almacenar clientes durante la ejecución del programa.
- Implementar las interfaces definidas en la capa usecase.
- Actuar como adaptador entre la lógica de negocio y la infraestructura.

Este repositorio NO contiene lógica de negocio.
Solo guarda, recupera y lista datos.
*/
type CustomerRepo struct {
	// byID almacena los clientes usando su ID como clave.
	// Ejemplo: byID[10] = Customer{ID:10, Name:"Juan", Email:"..."}
	byID map[int]domain.Customer
}

/*
NewCustomerRepo crea una nueva instancia del repositorio de clientes.

Se inicializa el mapa interno para evitar nil pointers
y permitir inserciones seguras desde el primer uso.
*/
func NewCustomerRepo() *CustomerRepo {
	return &CustomerRepo{
		byID: make(map[int]domain.Customer),
	}
}

/*
Create guarda un nuevo cliente en el repositorio.

Reglas:
- No permite IDs duplicados.
- La validación de campos (nombre, email, etc.)
  NO se hace aquí, sino en la capa usecase/domain.

Devuelve error si el ID ya existe.
*/
func (r *CustomerRepo) Create(c domain.Customer) error {
	if _, exists := r.byID[c.ID]; exists {
		return domain.ErrInvalidCustomerID
	}
	r.byID[c.ID] = c
	return nil
}

/*
List devuelve todos los clientes registrados.

Se retorna un slice para evitar exponer
la estructura interna del mapa.
*/
func (r *CustomerRepo) List() []domain.Customer {
	out := make([]domain.Customer, 0, len(r.byID))
	for _, c := range r.byID {
		out = append(out, c)
	}
	return out
}

/*
GetByID busca y devuelve un cliente por su ID.

Este método es CLAVE para el checkout:
- Permite obtener el nombre del cliente
- Evita que la CLI tenga que conocer la estructura interna

Devuelve error si el cliente no existe.
*/
func (r *CustomerRepo) GetByID(id int) (domain.Customer, error) {
	c, exists := r.byID[id]
	if !exists {
		return domain.Customer{}, domain.ErrInvalidCustomerID
	}
	return c, nil
}
