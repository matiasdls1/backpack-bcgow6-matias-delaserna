package main

import (
	"fmt"
	"os"
)

func almacenarProducto(id, cantidad int, precio float64) {
	text := fmt.Sprintf("%d,%.2f,%d", id, precio, cantidad)
	os.WriteFile("./ejercicio1.csv", []byte(text), 777)
}

func main() {
	id := 4500
	cantidad := 500
	precio := 3.555
	almacenarProducto(id, cantidad, precio)
}
