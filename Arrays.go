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
