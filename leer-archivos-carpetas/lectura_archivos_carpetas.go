package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Println("Elige una opción: ")
		fmt.Println("1. Ver archivos en la ruta actual")
		fmt.Println("2. Ver archivos y carpetas en la ruta anterior")
		fmt.Println("3. Ver archivos y carpetas en la ruta de la carpeta ~ (home)")
		fmt.Println("4. Salir")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Opción: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSuffix(choice, "\n")

		switch choice {
		case "1":
			cmd := exec.Command("ls", "-l")
			output, _ := cmd.CombinedOutput()
			fmt.Println("Archivos en la ruta actual:")
			fmt.Print(string(output))
		case "2":
			cmd := exec.Command("ls", "..", "-l")
			output, _ := cmd.CombinedOutput()
			fmt.Println("Archivos y carpetas en la ruta anterior:")
			fmt.Print(string(output))
		case "3":
			cmd := exec.Command("ls", os.Getenv("HOME"), "-l")
			output, _ := cmd.CombinedOutput()
			fmt.Println("Archivos y carpetas en la ruta de la carpeta ~ (home):")
			fmt.Print(string(output))
		case "4":
			fmt.Println("Adiós.")
			os.Exit(0)
		default:
			fmt.Println("Opción inválida. Inténtalo de nuevo.")
		}
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	for {
// 		fmt.Println("Elige una opción: ")
// 		fmt.Println("1. Ver archivos en la ruta actual")
// 		fmt.Println("2. Ver archivos y carpetas en la ruta anterior")
// 		fmt.Println("3. Ver archivos y carpetas en la ruta de la carpeta ~ (home)")
// 		fmt.Println("4. Salir")

// 		reader := bufio.NewReader(os.Stdin)
// 		fmt.Print("Opción: ")
// 		choice, _ := reader.ReadString('\n')

// 		switch choice {
// 		case "1\n":
// 			cmd := exec.Command("ls", "-l")
// 			output, _ := cmd.CombinedOutput()
// 			fmt.Println("Archivos en la ruta actual:")
// 			fmt.Print(string(output))
// 		case "2\n":
// 			cmd := exec.Command("ls", "..", "-l")
// 			output, _ := cmd.CombinedOutput()
// 			fmt.Println("Archivos y carpetas en la ruta anterior:")
// 			fmt.Print(string(output))
// 		case "3\n":
// 			cmd := exec.Command("ls", "~", "-l")
// 			output, _ := cmd.CombinedOutput()
// 			fmt.Println("Archivos y carpetas en la ruta de la carpeta ~ (home):")
// 			fmt.Print(string(output))
// 		case "4\n":
// 			fmt.Println("Adiós.")
// 			os.Exit(0)
// 		default:
// 			fmt.Println("Opción inválida. Inténtalo de nuevo.")
// 		}
// 	}
// }

// Este código hace uso de la biblioteca estándar de Go para ejecutar comandos externos (os/exec), y también hace uso de la biblioteca estándar para leer la entrada del usuario (bufio) y para acceder a las funciones del sistema operativo (os).

// El funcionamiento es similar al script de Bash anterior, con la diferencia de que en Go se usa una estructura de control switch para manejar las diferentes opciones elegidas por el usuario, en lugar de una estructura if-else. Además, en lugar de usar el comando cd, se ejecutan los comandos ls con los argumentos correspondientes para mostrar la información de los archivos y carpetas. La salida de estos comandos se muestra en pantalla usando
