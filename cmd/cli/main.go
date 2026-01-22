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

	for {
		fmt.Println("\n=== Sistema de Gestión de e-commerce (CLI) ===")
		fmt.Println("1) Productos")
		fmt.Println("0) Salir")
		fmt.Print("Opción: ")

		opcion := readLine(reader)

		switch opcion {
		case "1":
			productsMenu(reader, productRepo)
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
