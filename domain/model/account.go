package model

type Acconunt struct {
	Base      `valid:"required"`
	OwnerName string `json:"owner_name" valid:"notnull"`
}
