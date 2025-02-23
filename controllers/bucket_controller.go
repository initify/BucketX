package controllers

import (
	"bucketX/services"

	"github.com/gin-gonic/gin"
)

func ListBucketsController(c *gin.Context) {
	buckets, err := services.ListAllBuckets(c)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"buckets": buckets,
	})
}
