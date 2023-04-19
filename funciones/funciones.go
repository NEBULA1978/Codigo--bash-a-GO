package main

import "fmt"

func holaMundo() {
	fmt.Println("Hola mundo")
}

func parametros(nombre string, canal string) {
	fmt.Printf("Hola soy %s y me suscribo a %s\n", nombre, canal)
}

func main() {
	var nombre string
	var canal string

	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&nombre)

	fmt.Print("Ingrese su canal: ")
	fmt.Scanln(&canal)

	holaMundo()
	parametros(nombre, canal)
}

// En este ejemplo, la función holaMundo simplemente muestra un mensaje de "Hola mundo". La función parametros toma dos parámetros, nombre y canal, y los muestra en un mensaje. La función main es la función principal del programa y es donde se llaman a las funciones holaMundo y parametros.
