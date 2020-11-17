package main

import (
	"log"

	"github.com/EvgenyiK/aircrafts/controllers"
	"github.com/EvgenyiK/aircrafts/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r:= gin.Default()
	models.ConnectDatabase()
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	log.Fatal(r.Run(":8080"))
}