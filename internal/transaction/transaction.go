package transaction

import "github.com/jinzhu/gorm"

// the struct for the transaction servie, will implement the transaction service interface
type Service struct {
	DB *gorm.DB
}

// transaction structure
type Transaction struct {
	gorm.Model
	Customer string
	Product  string
}

// the interface for the transaction service
type TransactionService interface {
	GetTransaction(ID uint) (Transaction, error)
	GetTransactionsByCustomer(Customer string) ([]Transaction, error)
	PostTransaction(transaction Transaction) (Transaction, error)
	UpdateTransaction(ID uint, newTransaction Transaction) (Transaction, error)
	DeleteTransaction(ID uint) error
	GetAllTransactions() ([]Transaction, error)
}

//returns a new transaction service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

// retrieves transactions by their id from the db
func (s *Service) GetTransaction(ID uint) (Transaction, error) {
	var transaction Transaction
	if result := s.DB.First(&transaction, ID); result.Error != nil {
		return Transaction{}, result.Error
	}
	return transaction, nil
}

// retrieves all transactions by customer id
func (s *Service) GetTransactionsByCustomer(customer string) ([]Transaction, error) {
	var transactions []Transaction
	// check if this could be done with s.DB.Find(&transactions, ID)
	if result := s.DB.Find(&transactions).Where("Customer = ?", customer); result.Error != nil {
		return []Transaction{}, result.Error
	}
	return transactions, nil
}

// adds a new transaction to the db
func (s *Service) PostTransaction(transaction Transaction) (Transaction, error) {
	if result := s.DB.Save(&transaction); result.Error != nil {
		return Transaction{}, result.Error
	}
	return transaction, nil
}

// updates a transaction identified by ID with contents of newTransaction
func (s *Service) UpdateTransaction(ID uint, newTransaction Transaction) (Transaction, error) {
	transaction, err := s.GetTransaction(ID)
	if err != nil {
		return Transaction{}, err
	}
	if result := s.DB.Model(&transaction).Updates(newTransaction); result.Error != nil {
		return Transaction{}, result.Error
	}
	return transaction, nil
}

// deletes transaction identified by ID
func (s *Service) DeleteTransaction(ID uint) error {
	if result := s.DB.Delete(&Transaction{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// retrieves all transactions from the db
func (s *Service) GetAllTransactions() ([]Transaction, error) {
	var transactions []Transaction
	if result := s.DB.Find(&transactions); result.Error != nil {
		return []Transaction{}, result.Error
	}
	return transactions, nil
}
