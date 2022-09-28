package main

import "fmt"

/*
Ejercicio 2 - Ecommerce
Una importante empresa de ventas web necesita agregar una funcionalidad para agregar productos a los usuarios.
Para ello requieren que tanto los usuarios como los productos tengan la misma dirección de memoria en el main
del programa como en las funciones.
Se necesitan las estructuras:
Usuario: Nombre, Apellido, Correo, Productos (array de productos).
Producto: Nombre, precio, cantidad.
Se requieren las funciones:
Nuevo producto: recibe nombre y precio, y retorna un producto.
Agregar producto: recibe usuario, producto y cantidad, no retorna nada, agrega el producto al usuario.
Borrar productos: recibe un usuario, borra los productos del usuario.
*/

type usuario struct {
	nombre    string
	apellido  string
	correo    string
	productos []producto
}

type producto struct {
	nombre   string
	precio   float64
	cantidad int
}

type Producto interface {
	NuevoProducto(nombre string, precio float64) producto
}

type Usuario interface {
	AgregarProducto(usuario usuario, producto producto, cantidad int)
	BorrarProducto(usuario usuario)
}

func main() {
	usuario := usuario{
		nombre:   "Matias",
		apellido: "De la Serna",
		correo:   "matiasdelaserna@gmail",
		productos: []producto{
			{nombre: "lampara", precio: 1000, cantidad: 5},
			{nombre: "sillon", precio: 50000, cantidad: 2},
			{nombre: "monitor", precio: 20000, cantidad: 10},
		},
	}

	usuario.productos = append(usuario.productos, usuario.productos[0].NuevoProducto("mesa", 3))
	fmt.Println(usuario.productos)
}

func (p *producto) NuevoProducto(nombre string, precio float64) producto {
	producto := producto{
		nombre: nombre,
		precio: precio,
	}
	return producto
}

func (u *usuario) AgregarProducto(producto *producto, cantidad *int) {
	producto.cantidad = *cantidad
	u.productos = append(u.productos, *producto)
}

func (u *usuario) BorrarProducto(usuario usuario) {
	u.productos = []producto{}
}
