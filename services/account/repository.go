package account

import (
	"github.com/jamadeu/accounts/types"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) CreateAccount(account types.Account) error {
	return ar.db.Create(&account).Error
}
