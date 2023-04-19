package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Body  string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Mi primera página web en Go", "Bienvenido a mi primera página web en Go!"}
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, page)
}

func main() {
	http.HandleFunc("/", viewHandler)
	fmt.Println("El servidor está corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// Este ejemplo define una estructura Page que contiene un título y un cuerpo, y luego utiliza un manejador de vista para renderizar una plantilla HTML que muestra la página. La plantilla HTML puede verse así:

// <h1>{{.Title}}</h1>
// <p>{{.Body}}</p>
