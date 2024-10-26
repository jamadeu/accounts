package schemas

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"not null"`
	Document  string `gorm:"not null,unique"`
	Email     string `gorm:"not null,unique"`
	AccountID uint
}

type UserRepository interface {
	FindById(id uint) (*User, error)
	Create(user *User) (User, error)
}

type UserResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Document  string    `json:"document"`
	Email     string    `json:"email"`
	AccountID uint      `json:"accountId"`
}
