package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	for {
		fmt.Println("\033[1;36m MENU SCRIPT V.2 \033[0m")
		fmt.Println("1. Mostrar contenido del directorio actual")
		fmt.Println("2. Mostrar calendario del mes actual")
		fmt.Println("3. Mostrar fecha y hora actual")
		fmt.Println("4. Clonar repositorio de Github")
		fmt.Println("5. Pasar a la rama main")
		fmt.Println("6. Asegurarse de tener la última versión del código")
		fmt.Println("7. Crear una nueva rama")
		fmt.Println("8. Hacer cambios en el archivo")
		fmt.Println("9. Agregar cambios a la rama")
		fmt.Println("10. Subir cambios a Github")
		fmt.Println("11. Unir la nueva rama con la rama main")
		fmt.Println("12. Borrar una rama en Github")
		fmt.Println("13. Ver las ramas disponibles")
		fmt.Println("14. Borrar una rama local")
		fmt.Println("15. Eliminar archivo del control de versiones")
		fmt.Println("16. Salir")

		fmt.Print("Seleccione una opción: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]

		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Por favor, ingrese un número.")
			time.Sleep(1 * time.Second)
			continue
		}

		if option < 1 || option > 16 {
			fmt.Println("Opción inválida, por favor seleccione una opción válida.")
			time.Sleep(1 * time.Second)
			continue
		}

		switch option {
		case 1:
			fmt.Println("Mostrando contenido del directorio actual:")
			// En Go no hay un comando equivalente a "ls" en el shell
			// Aquí deberá utilizar la biblioteca de Go "os" para listar los archivos en el directorio actual
		case 2:
			fmt.Println("Mostrando calendario del mes actual:")
			// En Go no hay un comando equivalente a "cal", pero puede implementarse una función para mostrar el calendario.
		case 3:
			fmt.Println("Mostrando fecha y hora actual:")
			fmt.Println(time.Now().Format(time.RFC1123))
		case 4:
			fmt.Println("Clonando repositorio de Github...")
			fmt.Print("Introduce el nombre del repositorio a clonar: ")
			var repoName string
			fmt.Scanln(&repoName)
			// Aquí deberá implementar la lógica para clonar el repositorio de Github usando Go
		case 5:
			fmt.Println("Pasando a la rama main...")
			// Aquí deberá implementar la lógica para pasar a la rama main en Go
		case 6:
			fmt.Println("Asegurándose de tener la última versión del código...")
			// Aquí deberá implementar la lógica para asegurarse de tener la última versión del código en Go
		case 7:
			fmt.Println("Creando una nueva rama...")
			fmt.Print("Introduce el nombre de la nueva rama: ")
			var branchName string
			fmt.Scanln(&branchName)
			// Aquí deberá implementar la lógica para crear una nueva rama en Go
		case 8:
			fmt.Println("Haciendo cambios en el archivo...")
			// Aquí deberá implementar la lógica para hacer cambios en el archivo en Go
		case 9:
			fmt.Println("Agregando cambios a la rama...")
			// Aquí deberá implementar la lógica para agregar cambios a la rama en Go
		case 10:
			fmt.Println("Subiendo cambios a Github...")
			// Aquí deberá implementar la lógica para subir cambios a Github en Go
		case 11:
			fmt.Println("Uniendo la nueva rama con la rama main...")
			// Aquí deberá implementar la lógica para unir la nueva rama con la rama main en Go
		case 12:
			fmt.Println("Borrando una rama en Github...")
			fmt.Print("Introduce el nombre de la rama a borrar: ")
			var branchName string
			fmt.Scanln(&branchName)
			// Aquí deberá implementar la lógica para borrar una rama en Github en Go
		case 13:
			fmt.Println("Viendo las ramas disponibles...")
			// Aquí deberá implementar la lógica para ver las ramas disponibles en Go
		case 14:
			fmt.Println("Borrando una rama local...")
			fmt.Print("Introduce el nombre de la rama a borrar: ")
			var branchName string
			fmt.Scanln(&branchName)
			// Aquí deberá implementar la lógica para borrar una rama local en Go
		case 15:
			fmt.Println("Eliminando archivo del control de versiones...")
			fmt.Print("Introduce el nombre del archivo a eliminar: ")
			var fileName string
			fmt.Scanln(&fileName)
			// Aquí deberá implementar la lógica para eliminar un archivo del control de versiones en Go
		case 16:
			fmt.Println("Saliendo del script...")
			return
		}
		time.Sleep(1 * time.Second)
	}
}
