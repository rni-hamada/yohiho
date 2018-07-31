package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rni-hamada/yohiho/app/controllers"
	"github.com/rni-hamada/yohiho/app/db"
	"github.com/rni-hamada/yohiho/app/models"
)

func main() {
	db.Connect()

	initializeData()

	r := gin.Default()
	r.GET("/user", controllers.User)
	r.GET("/worker", controllers.Worker)
	r.GET("/worker/gacha", controllers.Gacha)

	r.GET("/competition", controllers.Competition)
	r.POST("/competition/:competition_id/participate", controllers.Participate)

	r.Run()
}

func initializeData() {
	db := db.Handler

	var count uint
	db.Table("users").Count(&count)

	log.Println("detected", count, "users")

	if count == 0 {
		user := models.User{Name: "yohihoer"}
		db.Create(&user)
	}
}
