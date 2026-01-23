package main

import (
	"bufio"   // Permite leer entradas del usuario desde la consola de forma eficiente
	"fmt"     // Proporciona funciones para imprimir texto en consola
	"os"      // Acceso a stdin/stdout y utilidades del sistema
	"strconv" // Conversión de strings a tipos numéricos
	"strings" // Manipulación de strings (trim, limpieza de saltos de línea)

	// Adaptadores: implementaciones concretas de repositorios en memoria.
	// Representan la capa de infraestructura.
	"github.com/aguirrethub/s-gestion-ecommerce/internal/adapters/memory"

	// Domain: entidades del negocio y reglas básicas (Product, Customer, Cart, errores).
	"github.com/aguirrethub/s-gestion-ecommerce/internal/domain"

	// Usecase: lógica de aplicación (casos de uso).
	// Orquesta las reglas de negocio sin depender de la infraestructura.
	"github.com/aguirrethub/s-gestion-ecommerce/internal/usecase"
)

/*
Función principal del programa.

Responsabilidades:
- Inicializar dependencias (repositorios).
- Mostrar el menú principal.
- Redirigir al usuario a los distintos submenús.
*/
func main() {
	// Reader central para toda la CLI.
	// Se reutiliza en todo el programa para leer entradas del usuario.
	reader := bufio.NewReader(os.Stdin)

	// Inicialización de repositorios en memoria.
	// Estos repositorios implementan interfaces definidas en la capa usecase,
	// lo que permite desacoplar la lógica del almacenamiento.
	productRepo := memory.NewProductRepo()
	customerRepo := memory.NewCustomerRepo()
	cartRepo := memory.NewCartRepo()

	// Bucle principal del sistema.
	// Se ejecuta indefinidamente hasta que el usuario elija salir.
	for {
		fmt.Println("\n=== Sistema de Gestión de e-commerce (CLI) ===")
		fmt.Println("1) Productos")
		fmt.Println("2) Clientes")
		fmt.Println("3) Carrito")
		fmt.Println("0) Salir")
		fmt.Print("Opción: ")

		// Se lee la opción como string para evitar errores de parseo directo.
		opcion := readLine(reader)

		// Enrutador del menú principal.
		switch opcion {
		case "1":
			productsMenu(reader, productRepo)

		case "2":
			customersMenu(reader, customerRepo)

		case "3":
			// El carrito necesita acceso a:
			// - CartRepository (carrito del cliente)
			// - ProductRepositoryForCart (validación de productos y stock)
			// - CustomerRepositoryForCheckout (obtener nombre del cliente)
			cartMenu(reader, cartRepo, productRepo, customerRepo)

		case "0":
			fmt.Println("Saliendo del sistema...")
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}

/*
productsMenu gestiona todas las operaciones relacionadas con productos.

Recibe:
- reader: para leer entradas del usuario.
- repo: interfaz ProductRepository (no depende de memory directamente).
*/
func productsMenu(reader *bufio.Reader, repo usecase.ProductRepository) {
	for {
		fmt.Println("\n--- Productos ---")
		fmt.Println("1) Crear producto")
		fmt.Println("2) Listar productos")
		fmt.Println("0) Volver")
		fmt.Print("Opción: ")

		op := readLine(reader)

		switch op {
		case "1":
			// Construcción de la entidad Product desde los datos ingresados.
			// La validación se realiza en el caso de uso.
			p := domain.Product{
				ID:    readInt(reader, "ID: "),
				Name:  readString(reader, "Nombre: "),
				Price: readFloat(reader, "Precio: "),
				Stock: readInt(reader, "Stock: "),
			}

			// Caso de uso: crea el producto aplicando reglas de negocio.
			if err := usecase.CreateProduct(repo, p); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Producto creado correctamente.")

		case "2":
			// Caso de uso: obtiene la lista de productos.
			products := usecase.ListProducts(repo)
			if len(products) == 0 {
				fmt.Println("No hay productos registrados.")
				continue
			}

			// Presentación en consola.
			for _, p := range products {
				fmt.Printf("ID:%d | %s | $%.2f | Stock:%d\n",
					p.ID, p.Name, p.Price, p.Stock)
			}

		case "0":
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}

/*
customersMenu gestiona las operaciones relacionadas con clientes.

Mantiene el mismo patrón que productos:
- La CLI solo captura datos.
- La lógica se delega a la capa usecase.
*/
func customersMenu(reader *bufio.Reader, repo usecase.CustomerRepository) {
	for {
		fmt.Println("\n--- Clientes ---")
		fmt.Println("1) Crear cliente")
		fmt.Println("2) Listar clientes")
		fmt.Println("0) Volver")
		fmt.Print("Opción: ")

		op := readLine(reader)

		switch op {
		case "1":
			c := domain.Customer{
				ID:    readInt(reader, "ID: "),
				Name:  readString(reader, "Nombre: "),
				Email: readString(reader, "Email: "),
			}

			if err := usecase.CreateCustomer(repo, c); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Cliente creado correctamente.")

		case "2":
			customers := usecase.ListCustomers(repo)
			if len(customers) == 0 {
				fmt.Println("No hay clientes registrados.")
				continue
			}

			for _, c := range customers {
				fmt.Printf("ID:%d | %s | %s\n",
					c.ID, c.Name, c.Email)
			}

		case "0":
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}

/*
cartMenu gestiona el carrito de compras y el proceso de checkout.

Recibe los repositorios necesarios para:
- Manipular el carrito
- Validar productos y stock
- Obtener información del cliente
*/
func cartMenu(
	reader *bufio.Reader,
	cartRepo usecase.CartRepository,
	productRepo usecase.ProductRepositoryForCart,
	customerRepo usecase.CustomerRepositoryForCheckout,
) {
	// Identificación del cliente que usará el carrito.
	customerID := readInt(reader, "CustomerID: ")

	for {
		fmt.Println("\n--- Carrito ---")
		fmt.Println("1) Ver carrito")
		fmt.Println("2) Agregar producto al carrito")
		fmt.Println("3) Quitar producto del carrito")
		fmt.Println("4) Vaciar carrito")
		fmt.Println("5) Total")
		fmt.Println("6) Checkout (Pagar)")
		fmt.Println("0) Volver")
		fmt.Print("Opción: ")

		op := readLine(reader)

		switch op {
		case "1":
			cart := usecase.ViewCart(cartRepo, customerID)
			if len(cart.Items) == 0 {
				fmt.Println("Carrito vacío.")
				continue
			}

			for _, it := range cart.Items {
				fmt.Printf(
					"ProdID:%d | %s | $%.2f | Cant:%d | Subtotal:$%.2f\n",
					it.ProductID, it.Name, it.Price,
					it.Quantity, it.Price*float64(it.Quantity),
				)
			}

			fmt.Printf("TOTAL: $%.2f\n",
				usecase.CartTotal(cartRepo, customerID))

		case "2":
			productID := readInt(reader, "ProductID: ")
			qty := readInt(reader, "Cantidad: ")

			if _, err := usecase.AddProductToCart(
				cartRepo, productRepo, customerID, productID, qty); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Producto agregado al carrito.")

		case "3":
			productID := readInt(reader, "ProductID a quitar: ")
			usecase.RemoveProductFromCart(cartRepo, customerID, productID)
			fmt.Println("Producto quitado.")

		case "4":
			usecase.ClearCart(cartRepo, customerID)
			fmt.Println("Carrito vaciado.")

		case "5":
			fmt.Printf("TOTAL: $%.2f\n",
				usecase.CartTotal(cartRepo, customerID))

		case "6":
			// Checkout: confirma la compra y genera comprobante.
			order, err := usecase.Checkout(
				cartRepo, productRepo, customerRepo, customerID)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			// Impresión del comprobante de pago.
			fmt.Println("\n=== COMPROBANTE DE PAGO ===")
			fmt.Println("Orden:", order.ID)
			fmt.Println("Cliente:", order.CustomerName, "(ID:", order.CustomerID, ")")
			fmt.Println("Fecha:", order.CreatedAt.Format("02-01-2006 15:04:05"))
			fmt.Println("--------------------------------------------------")
			fmt.Println("DETALLE:")

			for _, it := range order.Items {
				fmt.Printf(
					"ProdID:%d | %-15s | Unit:$%6.2f | Cant:%3d | Subtotal:$%7.2f\n",
					it.ProductID,
					it.Name,
					it.UnitPrice,
					it.Quantity,
					it.LineTotal,
				)
			}

			fmt.Println("--------------------------------------------------")
			fmt.Printf("TOTAL PAGADO: $%.2f\n", order.Total)
			fmt.Println("==================================================")

		case "0":
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}

/*
Funciones helper de entrada.

Centralizan la lectura y validación básica de datos,
evitando duplicación de código y errores de entrada.
*/

// Lee una línea completa y elimina espacios y saltos de línea.
func readLine(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

// Solicita y devuelve un string.
func readString(r *bufio.Reader, label string) string {
	fmt.Print(label)
	return readLine(r)
}

// Solicita un entero y repite hasta que sea válido.
func readInt(r *bufio.Reader, label string) int {
	for {
		fmt.Print(label)
		n, err := strconv.Atoi(readLine(r))
		if err == nil {
			return n
		}
		fmt.Println("Ingresa un número entero válido.")
	}
}

// Solicita un decimal y repite hasta que sea válido.
func readFloat(r *bufio.Reader, label string) float64 {
	for {
		fmt.Print(label)
		n, err := strconv.ParseFloat(readLine(r), 64)
		if err == nil {
			return n
		}
		fmt.Println("Ingresa un número decimal válido.")
	}
}
