package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKeyRepositoryInterface interface {
	RegisterKey(pixKey *PixKey) (*PixKey, error)
	FindKeyByKing(key string, kind string) (*PixKey, error)
	AddBank(bank *Banck) error
	AddAccount(acconunt *Acconunt) error
	FindAccount(id string) (*Acconunt, error)
}

type Acconunt struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" valid:"notnull"`
	Bank      *Banck    `valid:"-"`
	Number    string    `json:"number" valid:"notnull"`
	PixKeys   []*PixKey `valid:"-"`
}

func (account *Acconunt) isValid() error {
	_, err := govalidator.ValidateStruct(account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Banck, number string, ownerName string) (*Acconunt, error) {
	acconunt := Acconunt{
		OwnerName: ownerName,
		Bank:      bank,
		Number:    number,
	}

	acconunt.ID = uuid.NewV4().String()
	acconunt.CreatedAt = time.Now()

	err := acconunt.isValid()
	if err != nil {
		return nil, err
	}
	return &acconunt, nil
}
