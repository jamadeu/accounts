package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Balance      float64       `gorm:"not null"`
	User         User          `gorm:"not null"`
	Transactions []Transaction `gorm:"not null"`
}

type AccountRepository interface {
	CreateAccount(account Account) error
}

type AccountResponse struct {
	ID           uint          `json:"id"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	DeletedAt    time.Time     `json:"deletedAt,omitempty"`
	Balance      float64       `json:"balance"`
	User         UserResponse  `json:"user"`
	Transactions []Transaction `json:"transactions"`
}
