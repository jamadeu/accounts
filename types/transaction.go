package types

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	Type      string
	AccountID uint
}
