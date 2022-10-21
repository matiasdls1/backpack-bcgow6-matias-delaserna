package transactions

import (
	"fmt"

	"github.com/matiasdls1/backpack-bcgow6-matias-delaserna/go-testing/clase2tm/pkg/store"
)

/*
Repositorio, debe tener el acceso a la variable guardada en memoria.
Se debe crear el archivo repository.go
Se debe crear la estructura de la entidad
Se deben crear las variables globales donde guardar las entidades
Se debe generar la interface Repository con todos sus métodos
Se debe generar la estructura repository
Se debe generar una función que devuelva el Repositorio
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..)
*/

type Transaction struct {
	ID       int     `json:"id"`
	Code     string  `json:"code"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Date     string  `json:"date"`
}

var txs []Transaction // para updates en memoria - va a ser descartado

// var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	LastID() (int, error)
	Update(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	Delete(id int) error
	UpdateCodeAmount(id int, code string, amount float64) (Transaction, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() (transactions []Transaction, err error) {
	err = r.db.Read(&transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *repository) Store(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	var txs []Transaction
	err := r.db.Read(&txs)
	if err != nil {
		return Transaction{}, err
	}
	tx := Transaction{id, code, currency, amount, sender, receiver, date}
	txs = append(txs, tx)
	if err := r.db.Write(txs); err != nil {
		return Transaction{}, err
	}
	return tx, nil
}

func (r *repository) LastID() (int, error) {
	var txs []Transaction
	err := r.db.Read(&txs)
	if err != nil {
		return 0, err
	}
	if len(txs) == 0 {
		return 0, nil
	}

	return txs[len(txs)-1].ID, nil
}

func (r *repository) Update(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	tx := Transaction{Code: code, Currency: currency, Amount: amount, Sender: sender, Receiver: receiver, Date: date}
	updated := false
	var txs []Transaction
	err := r.db.Read(&txs)
	if err != nil {
		return Transaction{}, err
	}
	for i := range txs {
		if txs[i].ID == id {
			tx.ID = id
			txs[i] = tx
			updated = true
			break
		}
	}
	if !updated {
		return Transaction{}, fmt.Errorf("Transaction %d not found", id)
	}
	return tx, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	var txs []Transaction
	err := r.db.Read(&txs)
	if err != nil {
		return err
	}
	for i := range txs {
		if txs[i].ID == id {
			index = i
			deleted = true
			break
		}
	}
	if !deleted {

		return fmt.Errorf("Transaction %d not found", id)
	}
	txs = append(txs[:index], txs[index+1:]...)
	return nil
}

func (r *repository) UpdateCodeAmount(id int, code string, amount float64) (Transaction, error) {
	var tx Transaction
	updated := false
	err := r.db.Read(&tx)
	if err != nil {
		return Transaction{}, err
	}
	if tx.ID == id {
		tx.Code = code
		tx.Amount = amount
		updated = true
	}
	// for i := range txs {
	// }
	if !updated {
		return Transaction{}, fmt.Errorf("Transaction %d not found", id)
	}
	return tx, nil

}
