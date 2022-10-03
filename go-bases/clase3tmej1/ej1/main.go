package main

import (
	"fmt"
	"os"
)

type producto struct {
	id       int
	cantidad int
	precio   float64
}

func almacenarProducto(p producto) {
	text := fmt.Sprintf("%d,%.2f,%d", p.id, p.precio, p.cantidad)
	os.WriteFile("./ejercicio1.csv", []byte(text), 777)
}

func main() {
	producto := producto{id: 1, cantidad: 1, precio: 1.0}
	almacenarProducto(producto)
}
