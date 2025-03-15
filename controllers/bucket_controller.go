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

func CreateBucketController(c *gin.Context) {
	bucketName := c.PostForm("bucket_name")
	if bucketName == "" {
		c.JSON(400, gin.H{
			"error": "bucket_name is required",
		})
		return
	}
	err := services.CreateBucket(bucketName)
	if err!= nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "bucket created successfully",
	})
}
