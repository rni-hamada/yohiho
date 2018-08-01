package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rni-hamada/yohiho/app/models"
)

func Competition(c *gin.Context) {
	competitions := models.FetchCompetitions()

	c.JSON(200, gin.H{
		"contents": competitions,
	})
}

func Participate(c *gin.Context) {
	c.JSON(200, gin.H{
		"contents": []gin.H{
			{}, {},
		},
	})
}
