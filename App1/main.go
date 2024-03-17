package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p style='font-size: 20px;'>¡Bienvenido a mi aplicación Utesianos!</p>")
	fmt.Fprintf(w, "<p style='font-size: 16px;'>Mi nombre es Raymond Estrella Arias 1-18-8353</p>")

	// Agregar el botón de color rojo con estilo flat y evento onclick
	fmt.Fprintf(w, "<button id='botonRojo' style='background-color: red; color: white; border: none; padding: 10px 20px; text-align: center; text-decoration: none; display: inline-block; font-size: 14px; margin: 4px 2px; cursor: pointer;'> ¿Que hora es? </button>")

	// Agregar script JavaScript para manejar el evento onclick
	fmt.Fprintf(w, "<script>")
	fmt.Fprintf(w, "document.getElementById('botonRojo').onclick = function() { var fecha = new Date(); alert('La hora actual es: ' + fecha.toLocaleTimeString()); };")
	fmt.Fprintf(w, "</script>")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
