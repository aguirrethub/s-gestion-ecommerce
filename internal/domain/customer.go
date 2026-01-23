package domain

/*
Customer representa a un cliente del sistema.

Es una entidad de dominio:
- Modela un concepto central del negocio.
- No depende de infraestructura (repositorios, CLI, base de datos).
- Contiene solo datos relevantes del cliente.
*/
type Customer struct {
	ID    int    // Identificador único del cliente
	Name  string // Nombre del cliente
	Email string // Correo electrónico del cliente
}

/*
ValidateCustomer valida las reglas básicas del dominio para un cliente.

Responsabilidad:
- Garantizar que un Customer tenga datos coherentes antes de ser usado
  por los casos de uso (crear, asociar a carrito, etc.).

Reglas aplicadas:
- El ID debe ser mayor que 0.
- El nombre no puede estar vacío.
- El email debe tener un formato mínimo válido.

Nota:
- Esta función NO guarda al cliente.
- Solo valida reglas de negocio.
*/
func ValidateCustomer(c Customer) error {
	if c.ID <= 0 {
		return ErrInvalidCustomerID
	}
	if c.Name == "" {
		return ErrEmptyCustomerName
	}
	if !isValidEmailBasic(c.Email) {
		return ErrInvalidEmail
	}
	return nil
}

/*
isValidEmailBasic valida de forma simple el formato de un email.

Características:
- Verifica que el string no esté vacío.
- Verifica que contenga al menos un '@' y un '.'.

Importante:
- No es una validación RFC completa.
- Es intencionalmente simple para un proyecto CLI/educativo.
- Validaciones más complejas podrían hacerse con regex o librerías externas,
  pero eso no es responsabilidad del dominio en este contexto.
*/
func isValidEmailBasic(email string) bool {
	if email == "" {
		return false
	}

	hasAt := false
	hasDot := false

	for _, ch := range email {
		if ch == '@' {
			hasAt = true
		}
		if ch == '.' {
			hasDot = true
		}
	}

	return hasAt && hasDot
}
