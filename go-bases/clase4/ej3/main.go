package main

/*
Ejercicio 3 - Impuestos de salario #3
Repite el proceso anterior, pero ahora implementando “fmt.Errorf()”, para que el mensaje de error reciba
por parámetro el valor de “salary” indicando que no alcanza el mínimo imponible (el mensaje mostrado por
consola deberá decir: “error: el mínimo imponible es de 150.000 y el salario ingresado es de: [salary]”,
siendo [salary] el valor de tipo int pasado por parámetro).
*/

import (
	"fmt"
)

func main() {
	var salary int = 1000
	if salary < 150000 {
		fmt.Print(fmt.Errorf("error: el mínimo imponible es de 150.000 y el salario ingresado es de: %v", salary))
	} else {
		fmt.Print("Debe pagar el impuesto")
	}
}
