package user

import (
	"github.com/jamadeu/accounts/schemas"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindById(id string) (*schemas.User, error) {
	user := schemas.User{}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *schemas.User) (schemas.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return schemas.User{}, err
	}
	return *user, nil
}

func (r *UserRepository) ListUsers() (*[]schemas.User, error) {
	users := []schemas.User{}
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}
