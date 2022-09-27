package main

import (
	"fmt"
	"os"
	"strings"
)

func leerProductos() string { //(id, cantidad int, precio float64) {
	data, err := os.ReadFile("../ej1/ejercicio1.csv")
	if err != nil {
		panic(err)
	}
	//fileScanner := bufio.NewScanner(data)

	return string(data)
}

func main() {
	data := leerProductos()
	d := strings.Split(data, ",")
	fmt.Println("ID\t\tPrecio\tCantidad")
	fmt.Printf("%s\t\t%s\t%s", d[0], d[1], d[2])
}
