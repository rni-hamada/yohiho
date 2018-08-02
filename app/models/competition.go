package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rni-hamada/yohiho/app/db"
)

type Competition struct {
	gorm.Model
	Name         string `gorm:"size:255"`
	OpenAt       time.Time
	Participants []Participant
}

func FetchCompetition(competitionID uint) Competition {
	var competition Competition
	db.Handler.First(&competition, competitionID)
	return competition
}

func FetchCompetitions() []Competition {
	var competitions []Competition
	db.Handler.Find(&competitions)
	return competitions
}

func FetchParticipants(competition *Competition) []Participant {
	var participants []Participant
	db.Handler.Model(competition).Related(&participants)
	return participants
}

func IsParticipating(competition *Competition, worker *Worker) bool {
	var participant uint
	db.Handler.Table("participants").Where("worker_id = ? AND competition_id = ?", worker.ID, competition.ID).Count(&participant)
	return participant == 1
}

func Participate(competition *Competition, worker *Worker) {
	var participant = Participant{
		WorkerID:      worker.ID,
		CompetitionID: competition.ID,
		Pattern:       "wx",
	}
	db.Handler.Create(&participant)
}
