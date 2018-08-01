package controllers

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func Goat(c *gin.Context) {
	imgbytes, err := ioutil.ReadFile(fmt.Sprintf("assets/image/goat%d.jpg", 1))
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	c.Data(200, "image/jpeg", imgbytes)
}
