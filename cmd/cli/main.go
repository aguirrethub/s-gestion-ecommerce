package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== Sistema de Gestión de e-commerce (CLI) ===")
		fmt.Println("1) Productos")
		fmt.Println("0) Salir")
		fmt.Print("Opción: ")

		opcion := readLine(reader)

		switch opcion {
		case "1":
			fmt.Println("Entraste al módulo Productos (pendiente).")
		case "0":
			fmt.Println("Saliendo...")
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
