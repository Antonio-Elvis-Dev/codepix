package model
import "github.com/asaskevich/govalidator"

type Acconunt struct {
	Base      `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
	Bank *Bank `valid:"-"`
	Number string `json:"number" valid:"notnull"`
}

func (account *Acconunt) isValid() error{
	_, err := govalidator.ValidateStruct(account)
	if err != nil{
		return err
	}
	return nil
}