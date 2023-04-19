package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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
		fmt.Println("4. Salir")

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

		if option < 1 || option > 4 {
			fmt.Println("Opción inválida, por favor seleccione una opción válida.")
			time.Sleep(1 * time.Second)
			continue
		}

		switch option {
		case 1:
			fmt.Println("Mostrando contenido del directorio actual:")
			// En Go no hay un comando equivalente a "ls" en el shell
			// Aquí deberá utilizar la biblioteca de Go "os" para listar los archivos en el directorio actual
			files, _ := ioutil.ReadDir(".")
			for _, file := range files {
				fmt.Println(file.Name())
			}
		case 2:
			fmt.Println("Mostrando calendario del mes actual:")
			// En Go no hay un comando equivalente a "cal" en el shell
			// Aquí deberá utilizar la biblioteca de Go "time" para calcular y mostrar el calendario del mes actual
			t := time.Now()
			year, month, _ := t.Date()
			fmt.Println(month, year)
		case 3:
			fmt.Println("Mostrando fecha y hora actual:")
			fmt.Println(time.Now())
		case 4:
			fmt.Println("Saliendo del script...")
			return
		}

		fmt.Print("Presione [Enter] para continuar...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
