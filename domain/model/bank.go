package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Banck struct {
	Base `valid:"required"`
	Code string `json:"code" valid:"notnull"`
	Name string `json:"name" valid:"notnull"`
}

func (banck *Banck) isValid() error {

	_, err := govalidator.ValidateStruct(banck)
	if err != nil {
		return err
	}
	return nil
}

func NewBanck(code string, name string) (*Banck, error) {
	banck := Banck{
		Code: code,
		Name: name,
	}

	banck.ID = uuid.NewV4().String()
	banck.CreatedAt = time.Now()

	err := banck.isValid()
	if err != nil {
		return nil, err
	}
	return &banck, nil
}
