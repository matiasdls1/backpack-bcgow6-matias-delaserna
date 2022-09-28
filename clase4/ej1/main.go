package main

import "fmt"

/*
Ejercicio 1 - Impuestos de salario #1
En tu función “main”, define una variable llamada “salary” y asignarle un valor de tipo “int”.
Crea un error personalizado con un struct que implemente “Error()” con el mensaje “error: el salario
ingresado no alcanza el mínimo imponible" y lánzalo en caso de que “salary” sea menor a 150.000.
Caso contrario, imprime por consola el mensaje “Debe pagar impuesto”.
*/

type error struct{}

func (e *error) Error() string {
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func main() {
	var salary int = 1000
	if salary < 150000 {
		var err error
		fmt.Println(err)
	} else {
		fmt.Println("Debe pagar impuesto")
	}
}
