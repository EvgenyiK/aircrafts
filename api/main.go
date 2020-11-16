package main

import (
	"log"
	"net/http"

	//routes "github.com/EvgenyiK/aircrafts/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router:= gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "evgen",
		})
	})
	log.Fatal(router.Run(":8080"))
}