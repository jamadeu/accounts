package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Type      string
	AccountID uint
}

type TransactionResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Type      string    `json:"type"`
	AccountID uint      `json:"accountId"`
}