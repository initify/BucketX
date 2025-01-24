package controllers

import (
	"net/http"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"
)

func ServeDocsController(c *gin.Context) {
	htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
		SpecURL: "http://localhost:8080/docs/swagger.json",
		CustomOptions: scalar.CustomOptions{
			PageTitle: "BucketX API Reference",
		},
		DarkMode: true,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate API reference",
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))
}
