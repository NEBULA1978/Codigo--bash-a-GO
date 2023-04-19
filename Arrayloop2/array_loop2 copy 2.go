package main

import (
	"fmt"
	"sort"
)

func main() {
	// Agregar elementos a un arreglo
	numeros := []int{1, 2, 3}
	numeros = append(numeros, 4)
	fmt.Println(numeros) // imprime "[1 2 3 4]"

	// Eliminar elementos de un arreglo
	numeros = []int{1, 2, 3, 4, 5}
	numeros = append(numeros[:2], numeros[3:]...)
	fmt.Println(numeros) // imprime "[1 2 4 5]"

	// Extraer un rango de elementos de un arreglo
	numeros = []int{1, 2, 3, 4, 5}
	fmt.Println(numeros[1:4]) // imprime "[2 3 4]"

	// Ordenar un arreglo
	numeros = []int{5, 1, 3, 2, 4}
	sort.Ints(numeros)
	fmt.Println(numeros) // imprime "[1 2 3 4 5]"

	sort.Sort(sort.Reverse(sort.IntSlice(numeros)))
	fmt.Println(numeros) // imprime "[5 4 3 2 1]"

	// Obtener el índice de un elemento en un arreglo
	numeros = []int{1, 2, 3, 4, 5}
	buscar := 4
	for i, v := range numeros {
		if v == buscar {
			fmt.Println("El elemento", buscar, "está en la posición", i)
			break
		}
	}
	// imprime "El elemento 4 está en la posición 3"

}

// Resultado por consola
// [1 2 3 4]
// [1 2 4 5]
// [2 3 4]
// [1 2 3 4 5]
// [5 4 3 2 1]
// El elemento 4 está en la posición 3
