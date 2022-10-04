package transactions

/*
Servicio, debe contener la lógica de nuestra aplicación.
Se debe crear el archivo service.go.
Se debe generar la interface Service con todos sus métodos.
Se debe generar la estructura service que contenga el repositorio.
Se debe generar una función que devuelva el Servicio.
Se deben implementar todos los métodos correspondientes a las operaciones a realizar (GetAll, Store, etc..).
*/

type Service interface {
	GetAll() ([]Transaction, error)
	Store(code, currency string, amount float64, sender, receiver, date string) (Transaction, error)
	Update(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() ([]Transaction, error) {
	transactions, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (s *service) Store(code, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return Transaction{}, err
	}
	lastID++
	transaction, err := s.repository.Store(lastID, code, currency, amount, sender, receiver, date)
	if err != nil {
		return Transaction{}, err
	}
	return transaction, nil
}

func (s *service) Update(id int, code, currency string, amount float64, sender, receiver, date string) (Transaction, error) {
	return s.repository.Update(id, code, currency, amount, sender, receiver, date)
}
