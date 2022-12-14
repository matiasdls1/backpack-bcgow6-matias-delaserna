package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 2 - Impuestos de salario #2
Haz lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,
se implemente “errors.New()”.
*/

func main() {
	var salary int = 1000
	if salary < 150000 {
		fmt.Print(errors.New("error: el salario ingresado no alcanza el mínimo imponible"))
	} else {
		fmt.Print("Debe pagar el impuesto")
	}
}
