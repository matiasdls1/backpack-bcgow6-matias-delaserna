package main

import (
	"errors"
	"fmt"
)

/*
Bonus Track -  Impuestos de salario #4
Vamos a hacer que nuestro programa sea un poco más complejo.
1) Desarrolla las funciones necesarias para permitir a la empresa calcular:
a)
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10% en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
El mismo deberá indicar “error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
b)
Calcular el medio aguinaldo correspondiente al trabajador
Fórmula de cálculo de aguinaldo:
[mejor salario del semestre] / 12 * [meses trabajados en el semestre].
La función que realice el cálculo deberá retornar más de un valor, incluyendo un error en caso de que se ingrese un número
negativo.

2) Desarrolla el código necesario para cumplir con las funcionalidades requeridas, utilizando “errors.New()”, “fmt.Errorf()” y
“errors.Unwrap()”. No olvides realizar las validaciones de los retornos de error en tu función “main()”.
*/

func SalarioMensual(horasTrabajadas int, valorHora float64) (salario float64, err error) {
	if horasTrabajadas < 80 {
		err = errors.New("error: el trabajador no puede haber trabajado menos de 80 hs mensuales")
	} else {
		salario = float64(horasTrabajadas) * valorHora
	}

	return
}

func MedioAguinaldo(mejorSalario float64, mesesTrabajados int) (aguinaldo float64, err error) {
	if mejorSalario < 0 || mesesTrabajados < 0 {
		err = errors.New("error: los valores no pueden ser menos 0")
		return
	}

	if mesesTrabajados > 6 {
		err = errors.New("error: no puede trbajar mas de 6 meses por semestre")
	}
	aguinaldo = (mejorSalario / 12) * float64(mesesTrabajados)
	return
}

func main() {
	var horasTrabajadas, mesesTrabajados int = 80, 5
	var valorHora float64 = 300
	salario, err := SalarioMensual(horasTrabajadas, valorHora)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("El salario mensual es: ", salario)
		if salario >= 150000 {
			impuesto := salario * 0.1
			fmt.Println("Salario mayor a $150.000, se le descontaran: ", impuesto)
		}
	}

	medioAguinaldo, err := MedioAguinaldo(salario, mesesTrabajados)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Medio aguinaldo es: ", medioAguinaldo)
	}
}
