package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type PixKey struct {
	Base       `valid:"required"`
	Kind       string    `json:"kind" valid:"notnull"`
	Key        string    `json:"key"  valid:"notnull"`
	AcconuntID string    `json:"acconunt_id"  valid:"notnull"`
	Acconunt   *Acconunt `valid:"-"`
	Status     string    `json:"status"  valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)
	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(kind string, account *Acconunt, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:     kind,
		Key:      key,
		Acconunt: account,
		Status:   "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()

	err := pixKey.isValid()
	if err != nil {
		return nil, err
	}
	return &pixKey, nil
}
