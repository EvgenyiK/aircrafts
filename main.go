package main

import (
	"log"

	"github.com/EvgenyiK/aircrafts/controllers"
	"github.com/EvgenyiK/aircrafts/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	models.Connect()
	r.LoadHTMLGlob("templates/*")
	r.GET("/aircrafts", controllers.FindAircrafts)
	log.Fatal(r.Run(":8080"))
}