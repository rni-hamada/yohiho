package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rni-hamada/yohiho/app/db"
)

type Worker struct {
	gorm.Model
	Name    string `gorm:"size:255"`
	UserID  uint
	Hp      uint
	Atk     uint
	Def     uint
	Exp     uint
	Pattern string `gorm:"size:255"`
}

func FetchWorker(workerID uint) Worker {
	var competition Worker
	db.Handler.First(&competition, workerID)
	return competition
}

func CreateWorker(worker *Worker) {
	db.Handler.Create(worker)
}
