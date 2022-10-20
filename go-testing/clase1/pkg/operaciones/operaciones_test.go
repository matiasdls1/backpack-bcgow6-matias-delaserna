package operaciones

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddGood(t *testing.T) {
	// Arrange: valores de entrada y que quiero que devualva.
	num1 := 10
	num2 := 5
	esperado := 15

	// Act: llamar a la funcion
	resultado, err := Add(num1, num2)

	// Assert: evaluar el resultado de la funcion
	// pkg testing
	// if resultado != esperado {
	// 	t.Errorf("El numero resultado: %d, es distinto del esperado: %d", resultado, esperado)
	// }

	// pkg testify
	assert.Equal(t, esperado, resultado, "El numero resultado: %d, es distinto del esperado: %d", resultado, esperado)
	assert.Nil(t, err) // Evalua que error sea nil: pass si es nil
}

func TestAddBad(t *testing.T) {
	// Arrange: valores de entrada y que quiero que devualva.
	num1 := 0
	num2 := 5
	errorEsperado := fmt.Sprintf("num1 no puede ser: %d", num1)

	// Act: llamar a la funcion
	_, err := Add(num1, num2)

	// Assert: evaluar el resultado de la funcion
	// pkg testing
	// if resultado != esperado {
	// 	t.Errorf("El numero resultado: %d, es distinto del esperado: %d", resultado, esperado)
	// }

	// pkg testify
	assert.NotNil(t, err)                       // Evalua que error no sea nil: pass si NO es nil
	assert.ErrorContains(t, err, errorEsperado) // Evalua que el error recibido sea el esperado
}

func TestSubGood(t *testing.T) {
	// Arrange: valores de entrada y que quiero que devualva.
	num1 := 10
	num2 := 5
	esperado := 5

	// Act: llamar a la funcion
	resultado, err := Sub(num1, num2)

	// Assert: evaluar el resultado de la funcion
	// pkg testing
	// if resultado != esperado {
	// 	t.Errorf("El numero resultado: %d, es distinto del esperado: %d", resultado, esperado)
	// }

	// pkg testify
	assert.Equal(t, esperado, resultado, "El numero resultado: %d, es distinto del esperado: %d", resultado, esperado)
	assert.Nil(t, err) // Evalua que error sea nil: pass si es nil
}

func TestSubBad(t *testing.T) {
	// Arrange: valores de entrada y que quiero que devualva.
	num1 := 10
	num2 := 50

	// Act: llamar a la funcion
	_, err := Sub(num1, num2)

	// Assert: evaluar el resultado de la funcion
	// pkg testing
	// if resultado != esperado {
	// 	t.Errorf("El numero resultado: %d, es distinto del esperado: %d", resultado, esperado)
	// }

	// pkg testify
	assert.NotNil(t, err) // Evalua que error no sea nil: pass si NO es nil

}

func TestDividirGood(t *testing.T) {
	// Arrange
	num := 10
	den := 5
	esperado := 2

	// Act
	resultado, err := Dividir(num, den)

	// Assert
	assert.Equal(t, resultado, esperado)
	assert.Nil(t, err)
}

func TestDividirBad(t *testing.T) {
	// Arrange
	num := 10
	den := 0

	// Act
	_, err := Dividir(num, den)

	// Assert
	assert.NotNil(t, err)
}
