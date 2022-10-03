package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Ejercicio 2 - Registrando clientes

El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
Los datos requeridos para registrar a un cliente son:

Legajo
Nombre y Apellido
DNI
Número de teléfono
Domicilio

OK
Tarea 1: El número de legajo debe ser asignado o generado por separado y en forma previa a la carga de los restantes gastos.
Desarrolla e implementa una función para generar un ID que luego utilizarás para asignarlo como valor a “Legajo”.
Si por algún motivo esta función retorna valor “nil”, debe generar un panic que interrumpa la ejecución y aborte.

OK
Tarea 2: Antes de registrar a un cliente, debes verificar si el mismo ya existe. Para ello, necesitas leer los datos de un
archivo .txt. En algún lugar de tu código, implementa la función para leer un archivo llamado “customers.txt”
(como en el ejercicio anterior, este archivo no existe, por lo que la función que intente leerlo devolverá un error).
Debes manipular adecuadamente ese error como hemos visto hasta aquí.

Ese error deberá:
1.-   generar un panic;
2.- lanzar por consola el mensaje: “error: el archivo indicado no fue encontrado o está dañado”, y continuar con la ejecución
del programa normalmente.

Tarea 3: Luego de intentar verificar si el cliente a registrar ya existe, desarrolla una función para validar que todos los
datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar, al menos, dos valores.
Uno de los valores retornados deberá ser de tipo error para el caso de que se ingrese por parámetro algún valor cero
(recuerda los valores cero de cada tipo de dato, ej: 0, “”, nil).

Tarea 4: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los siguientes mensajes:
“Fin de la ejecución”, “Se detectaron varios errores en tiempo de ejecución” y “No han quedado archivos abiertos” (en ese orden).
Utiliza defer para cumplir con este requerimiento.

Requerimientos generales:
Utiliza recover para recuperar el valor de los panics que puedan surgir (excepto en la tarea 1).
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error (por ejemplo las que
intenten leer archivos).
Genera algún error, personalizandolo a tu gusto, utilizando alguna de las funciones que GO provee para ello
(realiza también la validación pertinente para el caso de error retornado).
*/

type cliente struct {
	legajo    int
	nombre    string
	dni       int
	telefono  int
	domicilio string
}

func GenerarLegajo() (int, bool) {
	var legajo int
	fmt.Print("Ingrese su legajo: ")
	fmt.Scanln(&legajo)
	if legajo <= 0 {
		return 0, false
	}
	return legajo, true
}

func VerificarCliente(cliente cliente) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("error: el archivo indicado no fue encontrado o esta dañado")
		}
	}()

	data, err := os.ReadFile("customers.txt")

	if err != nil {
		panic(err)
	} else {
		clientes := strings.Split(string(data), "\n")
		for _, c := range clientes {
			if strings.Contains(c, string(cliente.legajo)) {
				fmt.Println("Cliente ya existente: encontrado en la base de datos")
			}
		}
		fmt.Println("Cliente no existe: se lo registrará en la base de datos")
	}
}

func main() {
	var legajoOk bool
	cliente := &cliente{
		nombre:    "Matias de la Serna",
		dni:       40536806,
		telefono:  012345,
		domicilio: "Palermo",
	}
	cliente.legajo, legajoOk = GenerarLegajo()
	if !legajoOk {
		panic("Legajo incorrecto. Abortando ejecución.")
	}

	VerificarCliente(*cliente)

	fmt.Println("Ejecucion finalizada")
}
