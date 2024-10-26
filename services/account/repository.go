package account

import (
	"github.com/jamadeu/accounts/schemas"

	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccount(account schemas.Account) error {
	return r.db.Create(&account).Error
}
