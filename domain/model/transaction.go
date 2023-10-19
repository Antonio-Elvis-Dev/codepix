package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

const (
	TransactionPending   string = "pending"
	TransactionCompleted string = "completed"
	TransactionConfirmed string = "confirmed"
	TransactionError     string = "error"
)

type TransactionRepositoryInterface interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	Find(id string) (*Transaction, error)
}

type Transactions struct {
	Transaction []Transaction
}
type Transaction struct {
	Base              `valid:"required"`
	AcconuntForm      *Acconunt `valid:"-"`
	Amount            float64   `josn:"amount" valid:"notnull"`
	PixKeyTo          *PixKey   `valid:"-"`
	Status            string    `josn:"status" valid:"notnull"`
	Description       string    `josn:"description" valid:"notnull"`
	CancelDescription string    `josn:"cancel_description" valid:"notnull"`
}

func (t *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(t)

	if t.Amount <= 0 {
		return errors.New("the amount must be greater than 0")
	}

	if t.Status != TransactionPending && t.Status != TransactionCompleted && t.Status != TransactionError {
		return errors.New("invalid status for the transaction")
	}

	if t.PixKeyTo.AcconuntID == t.AcconuntForm.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewTransaction(accountForm *Acconunt, amount float64, pixKeyTo *PixKey, description string) (*Transaction, error) {
	transaction := Transaction{
		AcconuntForm: accountForm,
		Amount:       amount,
		Description:  description,
		PixKeyTo:     pixKeyTo,
		Status:       TransactionPending,
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()

	if err != nil {
		return nil, err
	}
	return &transaction, nil
}
func (t *Transaction) Complete() error {
	t.Status = TransactionCompleted
	t.UpdatedAt = time.Now()
	err := t.isValid()

	return err
}
func (t *Transaction) Confirmed() error {
	t.Status = TransactionConfirmed
	t.UpdatedAt = time.Now()
	err := t.isValid()

	return err
}
func (t *Transaction) Cancel(description string) error {
	t.Status = TransactionError
	t.UpdatedAt = time.Now()
	t.Description = description
	err := t.isValid()

	return err
}
