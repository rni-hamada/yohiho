package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rni-hamada/yohiho/app/db"
)

type User struct {
	gorm.Model
	Name    string `gorm:"size:255"`
	Workers []Worker
}

func GetUser() User {
	var user User
	db.Handler.First(&user)
	return user
}

func FetchWorkers(user *User) []Worker {
	var workers []Worker
	db.Handler.Model(user).Related(&workers)
	return workers
}
