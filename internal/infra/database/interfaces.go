package database

import "github.com/hudsonlhmartins/api-with-golang/internal/entity"

type UserInterface interface {
	Create(use *entity.User) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
}

type ProductInterface interface {
	Create(product *entity.Product) (*entity.Product, error)
	FindAll(page, limit int, sort string) ([]entity.Product, error)
	FindById(id string) (*entity.Product, error)
	Update(product *entity.Product) (*entity.Product, error)
	Delete(id string) error
}
