package main

type tienda struct {
	ecommerce string
}

type producto struct {
	nombre string
	tipo   string
	precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar() float64
}

func nuevoProducto(tipo, nombre string, precio float64) producto {
	producto := producto{
		nombre: nombre,
		tipo:   tipo,
		precio: precio,
	}
	return producto
}

func nuevaTienda(ecommerce string) tienda {
	tienda := tienda{
		ecommerce: ecommerce,
	}
	return tienda
}
