package main

import (
	"gogin/database"
	"gogin/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDataBase()

	router.GET("/books", usecase.FindBooks)
	router.GET("/book/:id", usecase.FindABook)
	router.POST("/book", usecase.CreateBook)
	router.PATCH("/book/:id", usecase.UpdateBook)

	router.Run()
}
