package controllers

import "github.com/gin-gonic/gin"

func Competition(c *gin.Context) {
	c.JSON(200, gin.H{
		"contents": []gin.H{
			{}, {}, {},
		},
	})
}

func Participate(c *gin.Context) {
	c.JSON(200, gin.H{
		"contents": []gin.H{
			{}, {},
		},
	})
}
