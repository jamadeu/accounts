package types

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Balance     float64
	User        User
	Transaction []Transaction
}

type User struct {
	gorm.Model
	Name      string
	Document  string
	Email     string
	AccountID uint
}

type Transaction struct {
	gorm.Model
	Type      string
	AccountID uint
}
