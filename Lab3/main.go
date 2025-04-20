package main

import (
	"ExperimentosLabs/library"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Crear una nueva libreria
	libraryInstance := library.NewLibrary()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n--- Menú de la Librería ---")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Verificar disponibilidad de un libro")
		fmt.Println("3. Eliminar libro")
		fmt.Println("4. Salir")
		fmt.Print("Seleccione una opción: ")

		scanner.Scan()
		option := strings.TrimSpace(scanner.Text())

		switch option {
		case "1":
			fmt.Print("Ingrese el título del libro a agregar: ")
			scanner.Scan()
			bookTitle := strings.TrimSpace(scanner.Text())
			libraryInstance.AddBook(bookTitle)
			fmt.Printf("Libro '%s' agregado a la librería.\n", bookTitle)

		case "2":
			fmt.Print("Ingrese el título del libro a verificar: ")
			scanner.Scan()
			bookTitle := strings.TrimSpace(scanner.Text())
			if libraryInstance.IsBookAvailable(bookTitle) {
				fmt.Printf("El libro '%s' está disponible.\n", bookTitle)
			} else {
				fmt.Printf("El libro '%s' no está disponible.\n", bookTitle)
			}

		case "3":
			fmt.Print("Ingrese el título del libro a eliminar: ")
			scanner.Scan()
			bookTitle := strings.TrimSpace(scanner.Text())
			if libraryInstance.RemoveBook(bookTitle) {
				fmt.Printf("El libro '%s' ha sido eliminado.\n", bookTitle)
			} else {
				fmt.Printf("El libro '%s' no está disponible para eliminar.\n", bookTitle)
			}

		case "4":
			fmt.Println("Saliendo del programa. ¡Hasta luego!")
			return

		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}

	}
}
