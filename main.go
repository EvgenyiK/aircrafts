package main

import (
	"log"
	"net/http"

	"github.com/EvgenyiK/aircrafts/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router:= gin.Default()
	models.Connect()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "evgen",
		})
	})
	log.Fatal(router.Run(":8080"))
}