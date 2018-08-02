package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rni-hamada/yohiho/app/models"
)

// Competition returns all open competitions
func Competition(c *gin.Context) {
	competitions := models.FetchCompetitions()

	c.JSON(200, gin.H{
		"contents": competitions,
	})
}

type ParticipateQuery struct {
	CompetitionId uint `form:"competition_id"`
	WorkerId      uint `form:"worker_id"`
}

func Participate(c *gin.Context) {
	var participateQuery ParticipateQuery
	if c.Bind(&participateQuery) == nil {
		competition := models.FetchCompetition(participateQuery.CompetitionId)
		if competition.ID == 0 {
			c.JSON(200, gin.H{"result": "no competition"})
			return
		}

		worker := models.FetchWorker(participateQuery.WorkerId)
		if worker.ID == 0 {
			c.JSON(200, gin.H{"result": "no worker"})
			return
		}

		log.Println("c=", competition, "w=", worker)

		if models.IsParticipating(&competition, &worker) {
			c.JSON(200, gin.H{"result": "already participating"})
		} else {
			models.Participate(&competition, &worker)
			c.JSON(200, gin.H{"result": "accepted"})
		}
	} else {
		c.JSON(200, gin.H{"result": "ng"})
	}
}
