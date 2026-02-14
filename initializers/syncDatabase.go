package initializers

import (
	"github.com/Bharat/go-bookstore/pkg/models"
)

func SyncDatabase() {
	Connect()
	db := GetDB()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Book{})
}



