package database

import (
	"fmt"
	"testing"

	"github.com/hudsonlhmartins/api-with-golang/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}

	defer sqlDB.Close()

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	userDB := NewUser(db)

	_, err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
}

func TestFindByEmail(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	db.AutoMigrate(&entity.User{})
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("could not open db: %v", err)
	}
	defer sqlDB.Close()
	user, _ := entity.NewUser("John Doe", "j@j.com", "123456")
	userDB := NewUser(db)
	_, err = userDB.Create(user)
	userFound, err := userDB.FindByEmail("j@j.com")
	fmt.Printf("useFoundId: %v\n", userFound.ID)
	fmt.Printf("userId: %v\n", user.ID)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}
