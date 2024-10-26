package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string
	Document  string
	Email     string
	AccountID uint
}
