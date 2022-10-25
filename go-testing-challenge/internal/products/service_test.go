package products

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// No hace falta testear repository.
// type StubStorage struct {
// 	dataMock []Product
// 	// func was called?
// }

func TestIntegrationGetAllBySeller(t *testing.T) {
	// Arrange
	// mockStorage := &MockStorage{
	// 	dataMock:      products,
	// }
	repository := NewRepository()
	service := NewService(repository)

	// Act
	result, err := service.GetAllBySeller("FEX112AC")

	// Assert
	assert.Nil(t, err)
	assert.Contains(t, result, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})

}
