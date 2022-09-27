package main

import "fmt"

type Alumno struct {
	nombre   string
	apellido string
	dni      int
	fecha    string
}

func (alumno Alumno) detalle() {
	fmt.Printf("Nombre: %v Apellido: %v Dni: %v Fecha: %v", alumno.nombre, alumno.apellido, alumno.dni, alumno.fecha)
}

func main() {
	var alumno Alumno
	alumno.nombre = "Matias"
	alumno.apellido = "de la Serna"
	alumno.dni = 40536806
	alumno.fecha = "21/09/2022"
	alumno.detalle()
}
