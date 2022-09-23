package main

import "fmt"

func main() {
	ej1()
	ej2()
	ej3()
	ej4()
}

func ej1() {
	palabra := "Miercoles"
	var cant int

	for i := 0; i < len(palabra); i++ {
		cant++
		fmt.Printf("%c\n", palabra[i])
	}
	fmt.Printf("\nCantidad de letras: %v\n", cant)
}

func ej2() {
	var prestamoOk bool
	var edad int = 25
	var empleado bool = true
	var antiguedad float64 = 1.3
	var sueldo int = 150000

	if edad > 22 {
		fmt.Print("El cliente es mayor de 22 años.\n")

		if empleado {
			fmt.Print("El cliente se encuentra empleado.\n")

			if antiguedad > 1 {
				fmt.Print("El cliente posee más de 1 año de antiguedad.\n")
				prestamoOk = true

				if sueldo > 100000 {
					fmt.Print("Al cliente no se le cobraran intereses sobre el prestamo.\n")

				} else {
					fmt.Print("Al cliente se le cobraran intereses sobre el prestamo.\n")

				}

			} else {
				fmt.Print("El cliente no posee más de 1 año de antiguedad. Prestamo cancelado.\n")
			}
		} else {
			fmt.Print("El cliente no se encuentra empleado. Prestamo cancelado.\n")
			prestamoOk = false
		}
	} else {
		fmt.Print("El cliente no es mayor de 22 años. Prestamo cancelado.\n")
		prestamoOk = false
	}

	if prestamoOk {
		fmt.Print("\nEl prestamo fue aceptado.\n")
	} else {
		fmt.Print("\nEl prestamo no fue aceptado.\n")
	}

}

func ej3() {
	var num int = 7
	switch num {
	case 1:
		fmt.Print("\nEnero")
	case 2:
		fmt.Print("\nFebrero")
	case 3:
		fmt.Print("\nMarzo")
	case 4:
		fmt.Print("\nAbril")
	case 5:
		fmt.Print("\nMayo")
	case 6:
		fmt.Print("\nJunio")
	case 7:
		fmt.Print("\nJulio")
	case 8:
		fmt.Print("\nAgosto")
	case 9:
		fmt.Print("\nSeptiembre")
	case 10:
		fmt.Print("\nOctubre")
	case 11:
		fmt.Print("\nNoviembre")
	case 12:
		fmt.Print("\nDiciembre")
	default:
		fmt.Print("Numero de mes equivocado.")
	}
}

func ej4() {
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	var cantMayores21 int = 0

	fmt.Printf("La edad de Benjamin es: %v", employees["Benjamin"])

	employees["Federico"] = 25

	for nombre, edad := range employees {
		if edad > 21 {
			cantMayores21++
		}
		if nombre == "Pedro" {
			delete(employees, nombre)
		}
	}
	fmt.Printf("\nMayores a 21 años: %v", cantMayores21)
}
