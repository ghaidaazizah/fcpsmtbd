package repository

import (
	"a21hc3NpZ25tZW50/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	Add(user model.User) error
	CheckAvail(user model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (u *userRepository) Add(user model.User) error {
	result := u.db.Create(&user) 
	if result.Error != nil {
		return result.Error 
	}
	return nil
}

func (u *userRepository) CheckAvail(user model.User) error {
	var existingUser model.User
	result := u.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&existingUser)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("user not found") 
	}
	return result.Error 
}
