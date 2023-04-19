package main

import (
	"fmt"
	"os/exec"
)

func main() {
	scripts := []string{
		"array_loop.sh",
		"case.sh",
		"condicionales.sh",
		"ejecutar-script-aqui.sh",
		"funciones.sh",
		"lectura_archivos-carpetas.sh",
		"lectura_archivos-entrar-ccarpetas2.sh",
		"lectura_archivos.sh",
		"lectura_archivos2-while.sh",
		"lectura_archivos2-while2.sh",
		"lectura_archivos2-while3.sh",
		"lectura_archivos2-while4.sh",
		"lectura_archivos2.sh",
		"manipular-archivos.sh",
		"manipular-archivos2.sh",
		"menubase.sh",
		"Menu-Github",
		"menubase4-completo-while-case.sh",
		"menubase4-completo-while-case2.sh",
		"menubasegit-Github-inputs.sh",
		"menubasegit-Github.sh",
		"pasos-git.sh",
		"primer_script.sh",
		"script_interactivo.sh",
		"tabla-multiplicar-input.sh",
		"variables.sh",
	}

	for _, script := range scripts {
		fmt.Printf("Ejecutando el script: %s\n", script)
		out, err := exec.Command("bash", script).Output()
		if err != nil {
			fmt.Printf("Error al ejecutar el script %s: %s\n", script, err)
		} else {
			fmt.Printf("Salida del script %s: %s\n", script, out)
		}
		fmt.Printf("El script %s ha terminado\n", script)
	}
}

// Este programa utiliza el paquete os/exec para ejecutar los scripts en una shell bash. El bucle for itera a través de la lista de nombres de script y los ejecuta uno por uno, imprimiendo un mensaje antes y después de cada ejecución.
