package main

import (
	"fmt"
)

func main() {
	ejercicio1()
	ejercicio2()
	ejercicio3()
	ejercicio4()
}

func ejercicio1() {
	nombre := "Matias de la Serna"
	direccion := "Fray Justo Santa Maria de Oro 2000"

	fmt.Print("\nEjercicio 1\n")
	fmt.Printf("Nombre: %v\nDireccion: %v\n", nombre, direccion)
}

func ejercicio2() {
	var temperatura int
	var humedad int
	var presion float64

	temperatura = 17
	humedad = 48
	presion = 1.021

	fmt.Print("\nEjercicio 2\n")
	fmt.Printf("Temperatura: %vº\nHumedad: %v%%\nPresion: %vhPa\n", temperatura, humedad, presion)
}

func ejercicio3() {
	/*
		var 1nombre string
		var apellido string
		var int edad
		1apellido := 6
		var licencia_de_conducir = true
		var estatura de la persona int
		cantidadDeHijos := 2
	*/
	var nombre string
	var apellido string
	var edad int
	var licenciaDeConducir = true
	var estaturaDeLaPersona float64
	cantidadDeHijos := 2
	nombre = "Matias"
	apellido = "De la Serna"
	edad = 25
	estaturaDeLaPersona = 1.87

	fmt.Print("\nEjercicio 3\n")
	fmt.Printf("Nombre: %v\nApellido: %v\nEdad: %v\nLicencia: %v\nEstatura: %v\nCantidad de hijos: %v\n", nombre, apellido, edad, licenciaDeConducir, estaturaDeLaPersona, cantidadDeHijos)
}

func ejercicio4() {
	/*
		var apellido string = "Gomez"
		var edad int = "35"
		boolean := "false";
		var sueldo string = 45857.90
		var nombre string = "Julián"
	*/
	var apellido string = "Gomez"
	var edad int = 35
	boolean := false
	var sueldo float64 = 45857.90
	var nombre string = "Julián"

	fmt.Print("\nEjercicio 4\n")
	fmt.Printf("Nombre: %v\nApellido: %v\nEdad: %v\nBoolean: %v\nSueldo: %v", nombre, apellido, edad, boolean, sueldo)

}
