package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	data []Transaction
}

type MockStorage struct {
	BeforeUpdate  Transaction
	ReadWasCalled bool
}

func (d *StubStore) Read(data interface{}) error { //(transaction []Transaction) {
	castedData := data.(*[]Transaction)
	*castedData = d.data
	return nil
}

func (d *StubStore) Write(data interface{}) error {
	castedData := data.(Transaction)
	d.data = append(d.data, castedData)
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	castedData := data.(*Transaction)
	*castedData = m.BeforeUpdate
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
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

	myStub := &StubStore{data: txs}
	repository := NewRepository(myStub)

	// Act
	resultado, err := repository.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, txs, resultado)
}

func TestUpdateCodeAmount(t *testing.T) {
	// Arrange
	beforeUpdate := Transaction{
		ID:       0,
		Code:     "A000",
		Currency: "PESOS",
		Amount:   10000,
		Sender:   "Matias",
		Receiver: "Pedro",
		Date:     "19/10/2022",
	}
	afterUpdate := Transaction{
		ID:       0,
		Code:     "B000",
		Currency: "PESOS",
		Amount:   200000.0,
		Sender:   "Matias",
		Receiver: "Pedro",
		Date:     "19/10/2022",
	}
	newCode := "B000"
	newAmount := 200000.0

	myMock := &MockStorage{BeforeUpdate: beforeUpdate, ReadWasCalled: false}
	repository := NewRepository(myMock)

	// Act
	resultado, err := repository.UpdateCodeAmount(beforeUpdate.ID, newCode, newAmount)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, afterUpdate, resultado)
	assert.True(t, myMock.ReadWasCalled)

}