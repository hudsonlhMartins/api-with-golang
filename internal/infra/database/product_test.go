package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hudsonlhmartins/api-with-golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&entity.Product{})

	product, _ := entity.NewProduct("Product 1", 10.5)
	productDB := NewProduct(db)
	dataCreate, err := productDB.Create(product)
	if err != nil {
		t.Fatalf("could not create product: %v", err)
	}

	var productFound entity.Product
	err = db.First(&productFound, "id = ?", dataCreate.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, dataCreate.ID, productFound.ID)
	assert.Equal(t, dataCreate.Name, productFound.Name)
	assert.Equal(t, dataCreate.Price, productFound.Price)

}

func TestFindAllProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	for i := 1; i < 24; i++ {
		product, _ := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		productDB.Create(product)
	}

	products, err := productDB.FindAll(1, 10, "asc")

	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 10, len(products))
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.Nil(t, err)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 23", products[2].Name)
}

func TestFindProductByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	product, _ := entity.NewProduct("Product 1", 10.5)
	productDB.Create(product)

	productFound, err := productDB.FindById(product.ID.String())
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	product, _ := entity.NewProduct("Product 1", 10.5)
	productDB.Create(product)

	product.Name = "Product 2"
	product.Price = 20.5

	productUpdated, err := productDB.Update(product)
	assert.Nil(t, err)
	assert.Equal(t, product.ID, productUpdated.ID)
	assert.Equal(t, product.Name, productUpdated.Name)
	assert.Equal(t, product.Price, productUpdated.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	defer sqlDB.Close()

	db.AutoMigrate(&entity.Product{})

	productDB := NewProduct(db)

	product, _ := entity.NewProduct("Product 1", 10.5)
	productDB.Create(product)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	product, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, product)
}
