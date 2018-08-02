package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Participant struct {
	gorm.Model
	CompetitionID uint
	WorkerID      uint
	Pattern       string `gorm:"size:255"`
	OpenAt        time.Time
}
