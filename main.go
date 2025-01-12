package main

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the Go Gin API!",
		})
	})

	router.POST("/image", func(c *gin.Context) {
     req, _ := io.ReadAll(c.Request.Body)
		 println(string(req));
	});

	router.Run(":8080")
}
