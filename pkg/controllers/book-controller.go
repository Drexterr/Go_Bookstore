package controllers

import (
	"net/http"
	"strconv"

	"github.com/Bharat/go-bookstore/initializers"
	"github.com/Bharat/go-bookstore/pkg/models"
	"github.com/gin-gonic/gin"
)

var NewBook models.Book

func GetBook(c *gin.Context) {
	newBooks,err := models.GetAllBooks()
	if err!= nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve data"})
			return
	}
	c.JSON(http.StatusOK, newBooks)

}

func GetBookByID(c *gin.Context) {
	bookID := c.Param("id")
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID "})
		return
	}
	bookDetails, _ := models.GetBookByID(ID)
	c.JSON(http.StatusOK, bookDetails)
}

func DuplicateStore(c *gin.Context){
SourceIDraw := c.Param("sourceid")
NewIDraw := c.Param("newid")
 newID, _ := strconv.ParseInt(NewIDraw, 0, 0)
 sourceID, _ := strconv.ParseInt(SourceIDraw,0,0)

 err := models.DuplicateStore(sourceID, newID)
 if err != nil{
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Failed to Duplicate Store"})
		return 
 }
 c.JSON(http.StatusOK, gin.H{
		"message": "Store duplicated successfully"})	

}

func GetBookByStoreId(c *gin.Context) {
	store_id := c.Param("Store_id")
	ID, err := strconv.ParseInt(store_id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID "})
		return
	}
	newBooks, err := models.GetBooksByStoreID(ID)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{ "error":"Failed to retrieve data"})
		return
	}
	c.JSON(http.StatusOK, newBooks)
}

func GetPriceByID(c *gin.Context) {
	bookID := c.Param("id")
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalidID"})
		return
	}
	bookDetails, _ := models.GetBookByID(ID)
	c.JSON(http.StatusOK, bookDetails.Price)
}

func CreateBook(c *gin.Context) {
	CreateBook := models.Book{}
	if err := c.ShouldBindJSON(&CreateBook); err != nil {

		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	b, err := CreateBook.CreateBook()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create book"})
		return
	}
	c.JSON(http.StatusCreated, b)
}

func DeleteBook(c *gin.Context) {
	bookID := c.Param("id")
	ID, err := strconv.ParseInt(bookID, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid ID"})
		return
	}
	book := models.DeleteBook(ID)
	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
    bookID := c.Param("id")
    ID, err := strconv.ParseInt(bookID, 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
        return
    }

    var updateData models.Book
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
        return
    }

    bookDetails, _ := models.GetBookByID(ID)
    if bookDetails.ID == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }

    if updateData.Name != "" {
        bookDetails.Name = updateData.Name
    }
    if updateData.Author != "" {
        bookDetails.Author = updateData.Author
    }
    if updateData.Publication != "" {
        bookDetails.Publication = updateData.Publication
    }
    if updateData.Price != 0 {
        bookDetails.Price = updateData.Price
    }
    if updateData.Store_id != 0 {
        bookDetails.Store_id = updateData.Store_id
    }

    initializers.GetDB().Save(&bookDetails)

    c.JSON(http.StatusOK, bookDetails)
}
