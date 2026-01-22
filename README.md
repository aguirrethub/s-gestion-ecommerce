# Sistema de Gestión de e-commerce (CLI)

Este proyecto implementa un Sistema de Gestión de e-commerce desarrollado en Go como una aplicación de línea de comandos (CLI).

El objetivo es aplicar los conceptos básicos del lenguaje Go vistos en la Unidad 1, tales como sintaxis, condicionales, estructuras de control, funciones y manejo de paquetes, utilizando un enfoque de programación funcional.

## Alcance

Funcionalidades incluidas:
- Gestión de productos.
- Gestión de clientes.
- Carrito de compras.
- Generación y confirmación de pedidos.
- Interfaz por consola (CLI).

Funcionalidades no incluidas:
- Pagos reales.
- Envíos.
- Persistencia en base de datos.
- Interfaz gráfica o web.

## Estructura del proyecto

- cmd/cli: punto de entrada del programa (main).
- internal/domain: modelos y reglas del negocio.
- internal/usecase: casos de uso del sistema.
- internal/adapters/memory: almacenamiento en memoria.

## Requisitos

- Go 1.20 o superior
- Git

## Ejecución

```bash
go mod tidy
go run ./cmd/cli
