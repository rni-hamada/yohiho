package controllers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/rni-hamada/yohiho/app/db"
	"github.com/rni-hamada/yohiho/app/models"
)

func User(c *gin.Context) {
	c.JSON(200, gin.H{
		"contents": models.GetUser(),
	})
}

func Worker(c *gin.Context) {
	var workers []models.Worker
	user := models.GetUser()
	db.Handler.Model(&user).Related(&workers)

	c.JSON(200, gin.H{
		"contents": workers,
	})
}

func Gacha(c *gin.Context) {
	user := models.GetUser()
	name := c.Param("name")
	hp := uint(10 + rand.Int31n(10))
	atk := uint(3 + rand.Int31n(3))
	def := uint(3 + rand.Int31n(3))

	worker := models.Worker{
		UserID:  user.ID,
		Name:    name,
		Hp:      hp,
		Atk:     atk,
		Def:     def,
		Exp:     0,
		Pattern: "",
	}

	models.CreateWorker(&worker)

	c.JSON(200, gin.H{"contents": worker})
}
