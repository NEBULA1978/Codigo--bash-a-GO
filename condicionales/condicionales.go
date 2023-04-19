package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ingresa tu nombre: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Ingresa tu edad: ")
	edadStr, _ := reader.ReadString('\n')
	fmt.Print("Ingresa el año actual: ")
	yearStr, _ := reader.ReadString('\n')

	edad, _ := strconv.Atoi(edadStr[:len(edadStr)-1])
	year, _ := strconv.Atoi(yearStr[:len(yearStr)-1])

	fmt.Printf("Hola, %s, tu edad es %d años.\n", name[:len(name)-1], edad)
	fmt.Println("====================================")
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Printf("%s, no eres mayor de edad\n", name[:len(name)-1])
	}

	fmt.Println("=========================================")
	if edad >= 18 && year == 2022 {
		fmt.Printf("%s, puedes votar\n", name[:len(name)-1])
	} else {
		fmt.Println("no puedes votar")
	}
}

// Este programa en Go hace uso de la biblioteca bufio para leer la entrada del usuario y la biblioteca strconv para convertir las entradas de texto a números. El programa realiza las mismas comprobaciones condicionales que el script anterior, y produce los mismos resultados.
