package database

import (
	"github.com/hudsonlhmartins/api-with-golang/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		DB: db,
	}
}

func (u *User) Create(user *entity.User) (*entity.User, error) {
	err := u.DB.Create(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
