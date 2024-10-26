package types

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Balance     float64
	User        User
	Transaction []Transaction
}

type AccountRepository interface {
	CreateAccount(account Account) error
}
