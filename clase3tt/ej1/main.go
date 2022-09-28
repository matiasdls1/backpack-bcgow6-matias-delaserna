package main

import "fmt"

/*
Ejercicio 1 - Red social
Una empresa de redes sociales requiere implementar una estructura usuario con funciones que vayan agregando información
a la estructura. Para optimizar y ahorrar memoria requieren que la estructura de usuarios ocupe el mismo lugar en memoria
para el main del programa y para las funciones.
La estructura debe tener los campos: Nombre, Apellido, Edad, Correo y Contraseña
Y deben implementarse las funciones:
Cambiar nombre: me permite cambiar el nombre y apellido.
Cambiar edad: me permite cambiar la edad.
Cambiar correo: me permite cambiar el correo.
Cambiar contraseña: me permite cambiar la contraseña.
*/

type usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

type Usuario interface {
	CambiarNombre(nombre, apellido string) string
	CambiarEdad(edad int) int
	CambiarCorreo() string
	CambiarContraseña() string
}

func main() {
	usuario := usuario{
		nombre:     "Matias",
		apellido:   "De la Serna",
		edad:       25,
		correo:     "matiasdelaserna@gmail.com",
		contraseña: "12345",
	}
	nuevoNombre := "Pedro"
	nuevoApellido := "Pappo"
	nuevaEdad := 30
	nuevoCorreo := "matias.delaserna@mercadolibre.com"
	nuevaContraseña := "67890"
	fmt.Printf("%p\n", &usuario)
	usuario.CambiarNombre(nuevoNombre, nuevoApellido)
	usuario.CambiarEdad(nuevaEdad)
	usuario.CambiarCorreo(nuevoCorreo)
	usuario.CambiarContraseña(nuevaContraseña)
	fmt.Println(usuario)
}

func (usuario *usuario) CambiarNombre(nombre, apellido string) string {
	fmt.Printf("%p", usuario)
	usuario.nombre = nombre
	usuario.apellido = apellido
	nuevoNombre := nombre + " " + apellido
	return nuevoNombre
}

func (usuario *usuario) CambiarEdad(edad int) int {
	fmt.Printf("%p", usuario)
	usuario.edad = edad
	return edad
}

func (usuario *usuario) CambiarCorreo(correo string) string {
	fmt.Printf("%p", usuario)
	usuario.correo = correo
	return correo
}

func (usuario *usuario) CambiarContraseña(contraseña string) string {
	fmt.Printf("%p", usuario)
	usuario.contraseña = contraseña
	return contraseña
}
