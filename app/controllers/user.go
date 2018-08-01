package controllers

import (
	"math/rand"

	"github.com/gin-gonic/gin"
	"github.com/rni-hamada/yohiho/app/models"
)

func User(c *gin.Context) {
	c.JSON(200, models.GetUser())
}

func Worker(c *gin.Context) {
	user := models.GetUser()
	c.JSON(200, models.FetchWorkers(&user))
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

	c.JSON(200, worker)
}
