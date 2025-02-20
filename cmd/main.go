package main

import (
	"go_book_api/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	api.InitDB()
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.POST("/book", api.CreateBook)
	r.GET("/books", api.GetBooks)
	r.GET("/book/:id", api.GetBook)
	r.PUT("/book/:id", api.UpdateBook)
	r.DELETE("/book/:id", api.DeleteBook)

	r.Run(":8080")
}
