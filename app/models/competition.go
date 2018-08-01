package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/rni-hamada/yohiho/app/db"
)

type Competition struct {
	gorm.Model
	Name   string `gorm:"size:255"`
	OpenAt time.Time
}

func FetchCompetitions() []Competition {
	var competitions []Competition
	db.Handler.Find(&competitions)
	return competitions
}
