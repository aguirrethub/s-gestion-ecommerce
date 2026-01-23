package usecase

import "github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

/*
CustomerRepository define el contrato que necesita la capa de casos de uso
para trabajar con clientes.

Principio aplicado:
- La capa usecase depende de interfaces, no de implementaciones concretas.
- Los adapters (memory, db, etc.) implementan esta interfaz.
*/
type CustomerRepository interface {
	// Create persiste un nuevo cliente.
	Create(c domain.Customer) error

	// List devuelve todos los clientes registrados.
	List() []domain.Customer
}

/*
CreateCustomer es un caso de uso de comando (modifica estado).

Responsabilidad:
- Validar que el cliente cumpla las reglas del dominio.
- Persistir el cliente usando el repositorio.

Flujo:
1) Valida el cliente con reglas de dominio.
2) Si es válido, delega la persistencia al repositorio.

Nota importante:
- La CLI NO valida clientes.
- El repositorio NO valida reglas de negocio.
- Este caso de uso es el punto correcto para coordinar ambas cosas.
*/
func CreateCustomer(repo CustomerRepository, c domain.Customer) error {
	// Validación de dominio (ID, nombre, email).
	if err := domain.ValidateCustomer(c); err != nil {
		return err
	}

	// Persistencia delegada al repositorio.
	return repo.Create(c)
}

/*
ListCustomers es un caso de uso de consulta.

Responsabilidad:
- Obtener la lista de clientes desde el repositorio.
- No aplica reglas de negocio ni validaciones.

Este tipo de función es deliberadamente simple:
- consulta datos
- no modifica estado
*/
func ListCustomers(repo CustomerRepository) []domain.Customer {
	return repo.List()
}
