package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aguirrethub/s-gestion-ecommerce/internal/adapters/memory"
	"github.com/aguirrethub/s-gestion-ecommerce/internal/domain"
	"github.com/aguirrethub/s-gestion-ecommerce/internal/usecase"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	productRepo := memory.NewProductRepo()
	customerRepo := memory.NewCustomerRepo()
	cartRepo := memory.NewCartRepo()

	for {
		fmt.Println("\n=== Sistema de Gestión de e-commerce (CLI) ===")
		fmt.Println("1) Productos")
		fmt.Println("2) Clientes")
		fmt.Println("3) Carrito")
		fmt.Println("0) Salir")
		fmt.Print("Opción: ")

		opcion := readLine(reader)

		switch opcion {
		case "1":
			productsMenu(reader, productRepo)
		case "2":
			customersMenu(reader, customerRepo)
		case "3":
			cartMenu(reader, cartRepo, productRepo)
		case "0":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

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
			p := domain.Product{
				ID:    readInt(reader, "ID: "),
				Name:  readString(reader, "Nombre: "),
				Price: readFloat(reader, "Precio: "),
				Stock: readInt(reader, "Stock: "),
			}
			if err := usecase.CreateProduct(repo, p); err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Producto creado correctamente.")
		case "2":
			products := usecase.ListProducts(repo)
			if len(products) == 0 {
				fmt.Println("No hay productos registrados.")
				continue
			}
			for _, p := range products {
				fmt.Printf("ID:%d | %s | $%.2f | Stock:%d\n", p.ID, p.Name, p.Price, p.Stock)
			}
		case "0":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

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
				fmt.Printf("ID:%d | %s | %s\n", c.ID, c.Name, c.Email)
			}
		case "0":
			return
		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func cartMenu(reader *bufio.Reader, cartRepo usecase.CartRepository, productRepo usecase.ProductRepositoryForCart) {
	customerID := readInt(reader, "CustomerID: ")

	for {
		fmt.Println("\n--- Carrito ---")
		fmt.Println("1) Ver carrito")
		fmt.Println("2) Agregar producto al carrito")
		fmt.Println("3) Quitar producto del carrito")
		fmt.Println("4) Vaciar carrito")
		fmt.Println("5) Total")
		fmt.Println("0) Volver")
		fmt.Print("Opción: ")

		op := readLine(reader)

		switch op {
		case "1":
			cart := usecase.ViewCart(cartRepo, customerID)

			// OJO: esto asume que tu domain.Cart tiene un slice llamado Items.
			// Si se llama distinto (ej: Lines, Products, etc.), pegá domain/cart.go y lo ajusto.
			if len(cart.Items) == 0 {
				fmt.Println("Carrito vacío.")
				continue
			}
			for _, it := range cart.Items {
				fmt.Printf("ProdID:%d | %s | $%.2f | Cant:%d | Subtotal:$%.2f\n",
					it.ProductID, it.Name, it.Price, it.Quantity, it.Price*float64(it.Quantity))
			}
			fmt.Printf("TOTAL: $%.2f\n", usecase.CartTotal(cartRepo, customerID))

		case "2":
			productID := readInt(reader, "ProductID: ")
			qty := readInt(reader, "Cantidad: ")

			_, err := usecase.AddProductToCart(cartRepo, productRepo, customerID, productID, qty)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			fmt.Println("Producto agregado al carrito.")

		case "3":
			productID := readInt(reader, "ProductID a quitar: ")
			usecase.RemoveProductFromCart(cartRepo, customerID, productID)
			fmt.Println("Producto quitado (si existía).")

		case "4":
			usecase.ClearCart(cartRepo, customerID)
			fmt.Println("Carrito vaciado.")

		case "5":
			fmt.Printf("TOTAL: $%.2f\n", usecase.CartTotal(cartRepo, customerID))

		case "0":
			return

		default:
			fmt.Println("Opción inválida.")
		}
	}
}

func readLine(r *bufio.Reader) string {
	s, _ := r.ReadString('\n')
	return strings.TrimSpace(s)
}

func readString(r *bufio.Reader, label string) string {
	fmt.Print(label)
	return readLine(r)
}

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
