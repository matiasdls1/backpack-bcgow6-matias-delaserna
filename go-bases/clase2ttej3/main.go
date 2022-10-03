package main

const (
	pequeño string = "pequeño"
	mediano string = "mediano"
	grande  string = "grande"
)

type tienda struct {
	t         Ecommerce
	productos []producto
}

type producto struct {
	p      Producto
	nombre string
	tipo   string
	precio float64
}

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto) []Producto
}

func nuevoProducto(tipo, nombre string, precio float64) *producto {
	return &producto{nombre: nombre, tipo: tipo, precio: precio}
}

func nuevaTienda() Ecommerce {
	return &tienda{}
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
	case pequeño:
		return p.precio
	case mediano:
		mantenimiento := (p.precio * 3) / 100
		return p.precio + mantenimiento
	case grande:
		mantenimiento := (p.precio * 6) / 100
		return p.precio + mantenimiento + 2500
	default:
		return 0
	}
}
