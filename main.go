package main

import (
	"errors"
	models "gin-gorm/Models"
	"gin-gorm/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var router = gin.Default()

func main() {
	database.StartDB()

	getAllBook()
	getBookById()
	createBook()
	updateBook()
	deleteBook()
	router.Run(":4000")
}

func getAllBook() {
	router.GET("/books", func(c *gin.Context) {
		db := database.GetDB()
		var books []models.Book
		err := db.Find(&books).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "data collected successfuly",
			"datas":   books,
		})
	})
}

func getBookById() {
	router.GET("/book/:id", func(c *gin.Context) {
		db := database.GetDB()
		book := models.Book{}
		id := c.Param("id")
		result := db.First(&book, id)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data Not Found",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "data collected succesfuly",
				"datas":   book,
			})
		}

	})
}

func createBook() {
	router.POST("/book/create", func(c *gin.Context) {
		db := database.GetDB()
		var book models.Book

		if err := c.ShouldBindJSON((&book)); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err := db.Create(&book).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Data Created Successfuly",
			"datas":   book,
		})
	})

}

func updateBook() {
	router.PUT("/book/update/:id", func(c *gin.Context) {
		db := database.GetDB()
		id := c.Param("id")
		var book models.Book
		result := db.First(&book, id)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data Not Found / Error",
			})
		} else {
			if err := c.ShouldBindJSON(&book); err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}

			result = db.Save(&book)
			c.JSON(http.StatusOK, gin.H{
				"message": "Data updated Successfuly",
			})
		}

	})
}

func deleteBook() {
	router.DELETE("book/delete/:id", func(c *gin.Context) {
		db := database.GetDB()
		id := c.Param("id")
		book := models.Book{}
		err := db.Where("id=?", id).Delete(&book).Error
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Data Deleted succesfully",
			})
		}

	})
}
