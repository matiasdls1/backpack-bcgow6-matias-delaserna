package transactions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock      []Transaction
	ReadWasCalled bool
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	a := data.(*[]Transaction)
	*a = m.dataMock
	if len(m.dataMock) == 0 {
		return fmt.Errorf("no transactions found")
	}
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	a := data.([]Transaction)
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}

func TestServiceIntegrationGetAll(t *testing.T) {
	// Arrange
	t1 := Transaction{
		ID:       0,
		Code:     "A000",
		Currency: "PESOS",
		Amount:   10000,
		Sender:   "Matias",
		Receiver: "Pedro",
		Date:     "19/10/2022",
	}
	t2 := Transaction{
		ID:       1,
		Code:     "A001",
		Currency: "DOLARES",
		Amount:   100,
		Sender:   "Juan",
		Receiver: "Lucas",
		Date:     "20/10/2022",
	}
	txs = []Transaction{t1, t2}
	mockStorage := &MockStorage{
		dataMock:      txs,
		ReadWasCalled: false,
	}

	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)

}
func TestServiceIntegrationUpdate(t *testing.T) {
	// Arrange
	inicial := []Transaction{
		{
			ID:       0,
			Code:     "A000",
			Currency: "PESOS",
			Amount:   10000,
			Sender:   "Matias",
			Receiver: "Pedro",
			Date:     "19/10/2022",
		},
		{
			ID:       1,
			Code:     "A001",
			Currency: "DOLARES",
			Amount:   100,
			Sender:   "Juan",
			Receiver: "Lucas",
			Date:     "20/10/2022",
		},
	}
	tx := Transaction{
		ID:       1,
		Code:     "A001",
		Currency: "EUROS",
		Amount:   500,
		Sender:   "Juan",
		Receiver: "Lucas",
		Date:     "20/10/2022",
	}
	mockStorage := &MockStorage{
		dataMock:      inicial,
		ReadWasCalled: false,
	}

	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	result, err := service.Update(tx.ID, tx.Code, tx.Currency, tx.Amount, tx.Sender, tx.Receiver, tx.Date)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[1], result)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestServiceIntegrationUpdateBad(t *testing.T) {
	// Arrange
	inicial := []Transaction{}
	tx := Transaction{
		ID:       1,
		Code:     "A001",
		Currency: "EUROS",
		Amount:   500,
		Sender:   "Juan",
		Receiver: "Lucas",
		Date:     "20/10/2022",
	}
	mockStorage := &MockStorage{
		dataMock:      inicial,
		ReadWasCalled: false,
	}

	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	_, err := service.Update(tx.ID, tx.Code, tx.Currency, tx.Amount, tx.Sender, tx.Receiver, tx.Date)

	// Assert
	assert.NotNil(t, err)
}

func TestServiceIntegrationUpdateBadNotFound(t *testing.T) {
	// Arrange
	inicial := []Transaction{
		{
			ID:       0,
			Code:     "A001",
			Currency: "EUROS",
			Amount:   500,
			Sender:   "Juan",
			Receiver: "Lucas",
			Date:     "20/10/2022",
		},
	}
	tx := Transaction{
		ID:       1,
		Code:     "A001",
		Currency: "EUROS",
		Amount:   500,
		Sender:   "Juan",
		Receiver: "Lucas",
		Date:     "20/10/2022",
	}
	mockStorage := &MockStorage{
		dataMock:      inicial,
		ReadWasCalled: false,
	}

	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	_, err := service.Update(tx.ID, tx.Code, tx.Currency, tx.Amount, tx.Sender, tx.Receiver, tx.Date)

	// Assert
	assert.NotNil(t, err)
}

func TestDelete(t *testing.T) {
	// Arrange
	inicial := []Transaction{
		{
			ID:       0,
			Code:     "A000",
			Currency: "PESOS",
			Amount:   10000,
			Sender:   "Matias",
			Receiver: "Pedro",
			Date:     "19/10/2022",
		},
		{
			ID:       1,
			Code:     "A001",
			Currency: "DOLARES",
			Amount:   100,
			Sender:   "Juan",
			Receiver: "Lucas",
			Date:     "20/10/2022",
		},
	}

	mockStorage := &MockStorage{
		dataMock:      inicial,
		ReadWasCalled: false,
	}

	// Act
	repository := NewRepository(mockStorage)
	service := NewService(repository)
	err := service.Delete(1)

	// Assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestDeleteBad(t *testing.T) {
	// Arrange
	inicial := []Transaction{}
	mockStorage := &MockStorage{
		dataMock:      inicial,
		ReadWasCalled: false,
	}
	repository := NewRepository(mockStorage)
	service := NewService(repository)

	// Act
	err := service.Delete(10)

	// Assert
	assert.NotNil(t, err)
}

func TestDeleteBadNotFound(t *testing.T) {
	// Arrange
	inicial := []Transaction{
		{
			ID:       0,
			Code:     "A000",
			Currency: "PESOS",
			Amount:   10000,
			Sender:   "Matias",
			Receiver: "Pedro",
			Date:     "19/10/2022",
		},
		{
			ID:       1,
			Code:     "A001",
			Currency: "DOLARES",
			Amount:   100,
			Sender:   "Juan",
			Receiver: "Lucas",
			Date:     "20/10/2022",
		},
	}

	mockStorage := &MockStorage{
		dataMock:      inicial,
		ReadWasCalled: false,
	}

	// Act
	repository := NewRepository(mockStorage)
	service := NewService(repository)
	err := service.Delete(10)

	// Assert
	assert.NotNil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
}
