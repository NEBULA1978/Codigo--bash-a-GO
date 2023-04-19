package main

import (
	"fmt"
)

func main() {
	var opcion string
	fmt.Print("Ingrese una opción [a-b]: ")
	fmt.Scanln(&opcion)

	switch opcion {
	case "a":
		fmt.Println("Ha ingresado la opción A")
	case "b":
		fmt.Println("Ha ingresado la opción B")
	default:
		fmt.Println("Ha ingresado un valor fuera del rango o opción inválida")
	}
}
