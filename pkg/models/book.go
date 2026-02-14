package models

import (
	
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Price       int    `json:"price"`
	Store_id    int    `json:"store_id"`

}



func (b *Book) CreateBook() (*Book, error) {

	result := db.Create(&b)
	return b, result.Error
}

func GetAllBooks() ([]Book, error) {
	var Books []Book
	result := db.Find(&Books)
	return Books, result.Error
}

func DuplicateStore(sourceID int64, newID int64) error {
	query := `INSERT INTO books (name, author, publication, price, store_id, created_at, updated_at)
		SELECT name, author, publication, price, ? 
		FROM books 
		WHERE store_id = ? AND deleted_at IS NULL`
	err := db.Exec(query, newID, sourceID).Error
	return err
}

func GetBookByID(id int64) (*Book, error) {
	var GetBook Book
	result := db.First(&GetBook, id)
	return &GetBook, result.Error
}

func DeleteBook(id int64) error {
	var book Book
	result := db.Delete(&book, id)
	return result.Error
}

func GetPriceByID(id int64) (*Book, error) {
	var book Book
	result := db.Find(&book, id)
	return &book, result.Error
}

func GetBooksByStoreID(Store_id int64) ([]Book, error) {
	var Books []Book
	result := db.Where("store_id = ?", Store_id).Find(&Books)
	return Books, result.Error
}
