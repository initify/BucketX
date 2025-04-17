package controllers

import (
	"bucketX/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddNode(c *gin.Context) {
	var newNode services.Node
	if err := c.ShouldBindJSON(&newNode); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	err := services.AddNode(newNode)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "node added successfully"})
}
