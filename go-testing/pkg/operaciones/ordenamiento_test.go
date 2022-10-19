package operaciones

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdenarGood(t *testing.T) {
	// Arrange
	numeros := []int{8, 2, 6, 10, 4}
	esperado := []int{2, 4, 6, 8, 10}

	// Act
	resultado, err := Ordenar(numeros)

	// Assert
	assert.Equal(t, resultado, esperado)
	assert.Nil(t, err)
}

func TestOrdenarBad(t *testing.T) {
	// Arrange
	numeros := []int{}

	// Act
	_, err := Ordenar(numeros)

	// Assert
	assert.NotNil(t, err)
}
