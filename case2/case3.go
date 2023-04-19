package main

import (
	"bufio"
	"fmt"
	"os"
)

func switchCase(option string) {
	switch option {
	case "a":
		fmt.Println("Ha ingresado la opción A")
	case "b":
		fmt.Println("Ha ingresado la opción B")
	default:
		if option >= "c" && option <= "z" {
			fmt.Println("Ha ingresado un valor fuera del rango")
		} else {
			fmt.Println("Opción inválida")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingrese una opción [a-b]: ")
	option, _ := reader.ReadString('\n')
	switchCase(option[:len(option)-1])
}
