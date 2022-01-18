package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/max10rogerio/book-go-example/database"
	"github.com/max10rogerio/book-go-example/models"
)

func ShowBook(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "ID has to be integer",
			"error":   err.Error(),
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, id).Error

	if err != nil {
		context.JSON(404, gin.H{
			"message": "book not found",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func ShowBooks(context *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		context.JSON(400, gin.H{
			"message": "cannot list books",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(200, gin.H{
		"books": books,
	})
}

func CreateBook(context *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := context.ShouldBindJSON(&book)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "cannot bind json",
			"error":   err.Error(),
		})

		return
	}

	err = db.Create(&book).Error

	if err != nil {
		context.JSON(400, gin.H{
			"message": "cannot create book",
			"error":   err.Error(),
		})

		return
	}

	context.JSON(200, book)
}

func UpdateBook(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "ID has to be integer",
			"error":   err.Error(),
		})

		return
	}

	db := database.GetDatabase()

	var book models.Book

	err = db.First(&book, id).Error

	if err != nil {
		context.JSON(404, gin.H{
			"message": "book not found with id: " + strconv.Itoa(id),
			"error":   err.Error(),
		})

		return
	}

	err = context.ShouldBindJSON(&book)

	if err != nil {
		context.JSON(400, gin.H{
			"message": "cannot bind json",
			"error":   err.Error(),
		})

		return
	}

	err = db.Save(&book).Error

	if err != nil {
		context.JSON(400, gin.H{
			"message": "cannot update book",
			"error":   err.Error(),
		})
	}

	context.JSON(200, book)
}

func DeleteBook(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(400, gin.H{
			"message": "ID has to be integer",
			"error":   err.Error(),
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Book{}, id).Error

	if err != nil {
		context.JSON(400, gin.H{
			"message": "cannot delete book",
			"error":   err.Error(),
		})

		return
	}

	context.Status(204)
}
