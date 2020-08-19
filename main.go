package main

import (
	"gogin/controllers"
	"gogin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	models.ConnectDataBase()

	router.GET("/books", controllers.FindBooks)
	router.GET("/book/:id", controllers.FindABook)
	router.POST("/books", controllers.CreateBook)

	router.Run()
}
