package clase2

import (
	"fmt"
)

const (
	salario  = 200000
	minimo   = "minimo"
	maximo   = "maximo"
	promedio = "promedio"
)

func Practica1() {
	//fmt.Println(ejercicio1(salario))
	ejercicio4()
}

func ejercicio1(salario float64) float64 {
	if salario > 50000 && salario <= 150000 {
		return salario * 0.17
	} else if salario > 150000 {
		return salario * 0.1
	}

	return 0.0
}

func ejercicio2(notas ...int) (promedio float64, err error) {
	suma := 0
	if len(notas) != 0 {
		for _, value := range notas {
			if value >= 0 {
				suma += value
			} else {
				return 0, fmt.Errorf("%v es negativo", value)
			}
		}
		promedio := float64(suma) / float64(len(notas))
		return promedio, nil
	} else {
		return 0, fmt.Errorf("No se mandaron las notas")
	}
}

func ejercicio3(minutos int, categoria string) (salario float64) {
	var salarioPorHora float64
	var extra float64
	var horasTrabajadas float64
	switch categoria {
	case "A":
		salarioPorHora = 3000
		extra = 1.5
	case "B":
		salarioPorHora = 1500
		extra = 1.2
	case "C":
		salarioPorHora = 1000
		extra = 1
	}

	horasTrabajadas = float64(minutos) / 60
	salario = salarioPorHora * horasTrabajadas * extra
	return salario
}

func min(notas []int) float64 {
	min := notas[0]

	for _, value := range notas {
		if min > value {
			min = value
		}
	}

	return float64(min)
}

func max(notas []int) float64 {
	max := notas[0]

	for _, value := range notas {
		if max < value {
			max = value
		}
	}

	return float64(max)
}

func prom(notas []int) (promedio float64) {
	suma := 0
	if len(notas) > 0 {
		for _, value := range notas {
			suma += value
		}
		promedio = float64(suma) / float64(len(notas))
		return
	} else {
		return
	}
}

func orquestador(operacion string) (string, func([]int) float64) {
	switch operacion {
	case maximo:
		return maximo, max
	case minimo:
		return minimo, min
	case promedio:
		return promedio, prom
	}
	return "Operacion Invalida", nil
}

func ejercicio4() {
	notas := []int{2, 5, 8, 10, 7}
	operacionDevuelta, value := orquestador(maximo)

	fmt.Println(operacionDevuelta, value(notas))
}
