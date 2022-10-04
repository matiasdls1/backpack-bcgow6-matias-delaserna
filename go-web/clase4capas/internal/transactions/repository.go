package transactions

import "fmt"

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

var txs []Transaction
var lastID int

type Repository interface {
	GetAll() ([]Transaction, error)
	Store(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	LastID() (int, error)
	Update(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error)
}

type repository struct{} // struct implementa los metodos de la interfaz

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Transaction, error) {
	return txs, nil
}

func (r *repository) Store(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	tx := Transaction{id, code, currency, amount, sender, receiver, date}
	txs = append(txs, tx)
	lastID = tx.ID
	return tx, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Update(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	tx := Transaction{Code: code, Currency: currency, Amount: amount, Sender: sender, Receiver: receiver, Date: date}
	updated := false
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
