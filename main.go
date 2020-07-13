package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rayning0/go-blog/controllers"
	"github.com/rayning0/go-blog/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook) //only updates specified keys in a book
	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
