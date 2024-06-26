package database

import (
	"github.com/hudsonlhmartins/api-with-golang/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{
		DB: db,
	}
}

func (p *Product) Create(product *entity.Product) (*entity.Product, error) {
	if err := p.DB.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	if err := p.DB.First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p *Product) Update(product *entity.Product) (*entity.Product, error) {
	_, err := p.FindById(product.ID.String())
	/*
		Temos que da um FindById pos o save se o registro nao existir ele cria um novo,
		para evitar isso temos que verificar se o registro existe, e se ele não existir retornar um erro
	*/
	if err != nil {
		return nil, err
	}
	if err := p.DB.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error

	}
	return products, err
}

func (p *Product) Delete(id string) error {
	product, err := p.FindById(id)
	if err != nil {
		return err
	}
	return p.DB.Delete(product).Error
}
