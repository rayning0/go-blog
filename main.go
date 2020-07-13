package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rayning0/go-blog/controllers"
	"github.com/rayning0/go-blog/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/users", controllers.FindUsers)
	r.POST("/users", controllers.CreateUser)
	r.GET("/users/:id", controllers.FindUser)
	r.PATCH("/users/:id", controllers.UpdateUser) //only updates specified keys in a User
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.Run()
}
