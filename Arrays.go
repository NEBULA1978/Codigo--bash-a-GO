package main

import "fmt"

func main() {
	numeros := []int{1, 2, 3, 4, 5, 6, 7, 8}
	nombres := []string{"juan", "pedro", "luffy", "goku", "naruto"}
	rangos := append([]string{}, []string(make([]string, 27, 27))...)
	for i := 1; i <= 30; i++ {
		rangos = append(rangos, fmt.Sprintf("%d", i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		rangos = append(rangos, fmt.Sprintf("%c", i))
	}

	fmt.Println("===== IMPRIMIR  VALORES ===")

	fmt.Println(numeros)
	fmt.Println(nombres)
	fmt.Println(rangos)

	fmt.Println("== Tamaño arreglos ==")

	fmt.Println("Tamaño de numeros:", len(numeros))
	fmt.Println("Tamaño de nombres:", len(nombres))
	fmt.Println("Tamaño de rangos:", len(rangos))

	fmt.Println("=== Elemento de arreglo ===")

	fmt.Println("Elemento numero 3 del arreglo numeros:", numeros[2])
	fmt.Println("Elemento numero 4 del arreglo nombres:", nombres[3])
	fmt.Println("Elemento numero 28 del arreglo numeros:", rangos[27])

	fmt.Println("=== Manipular arreglos ===")

	numeros = append(numeros[:0], numeros[1:]...)

	fmt.Println("Numeros:", numeros)

	numeros = append([]int{1}, numeros...)

	fmt.Println("Numeros:", numeros)

	fmt.Println("===== ITERACION FOR ========")

	for _, num := range numeros {
		fmt.Println("Numero:", num)
	}

	fmt.Println("===============================")

	for i := 0; i < len(numeros); i++ {
		fmt.Println("Numero:", numeros[i])
	}
}

// CODIGO EN BASH

// # Autor: Ramón Pascual
// # Arreglo e Iteración

// numeros=(1 2 3 4 5 6 7 8) # elementos diferenciados por espacios
// nombres=("juan" "pedro" "luffy" "goku" "naruto")
// rangos=({A..Z} {1..30})

// echo "===== IMPRIMIR  VALORES ==="

// echo "${numeros[@]}"
// echo "${nombres[@]}"
// echo "${rangos[@]}"

// echo "== Tamaño arreglos =="

// echo "Tamaño de numeros: ${#numeros[@]}" # sintaxis para contar el número de elementos del arreglo
// echo "Tamaño de nombres: ${#nombres[@]}"
// echo "Tamaño de rangos: ${#rangos[*]}"

// echo "=== Elemento de arreglo ==="

// echo "Elemento numero 3 del arreglo numeros: ${numeros[2]}"
// echo "Elemento numero 4 del arreglo nombres: ${nombres[3]}"
// echo "Elemento numero 28 del arreglo numeros: ${rangos[27]}"

// echo "=== Manipular arreglos ==="

// unset 'numeros[0]' # elimina el elemento de esa posición

// echo "Numeros: ${numeros[*]}"

// numeros[0]=1

// echo "Numeros: ${numeros[*]}"

// echo "===== ITERACION FOR ========"

// for num in "${numeros[@]}"; do # for each
//     echo "Numero: $num"
// done

// echo "==============================="

// for ((i = 0; i < ${#numeros[@]}; i++)); do
//     echo "Numero: ${numeros[i]}"
// done

// # EQUIVALENTE EN PYTHON3

// # hastack-admiracion/usr/bin/env python3
// # # Author: Ramón Pascual
// # # Array e Iteracion

// # numeros = [1, 2, 3, 4, 5, 6, 7, 8]
// # nombres = ["juan", "pedro", "luffy", "goku", "naruto"]
// # rangos = list(map(str, list(range(1, 31)))) + [chr(i) for i in range(ord('A'), ord('Z')+1)]

// # print("===== IMPRIMIR  VALORES ===")

// # print(numeros)
// # print(nombres)
// # print(rangos)

// # print("== Tamaño arreglos ==")

// # print("Tamaño de numeros:", len(numeros))
// # print("Tamaño de nombres:", len(nombres))
// # print("Tamaño de rangos:", len(rangos))

// # print("=== Elemento de arreglo ===")

// # print("Elemento numero 3 del arreglo numeros:", numeros[2])
// # print("Elemento numero 4 del arreglo nombres:", nombres[3])
// # print("Elemento numero 28 del arreglo numeros:", rangos[27])

// # print("=== Manipular arreglos ===")

// # numeros.pop(0)

// # print("Numeros:", numeros)

// # numeros.insert(0, 1)

// # print("Numeros:", numeros)

// # print("===== ITERACION FOR ========")

// # for num in numeros:
// #     print("Numero:", num)

// # print("===============================")

// # for i in range(len(numeros)):
// #     print("Numero:", numeros[i])
