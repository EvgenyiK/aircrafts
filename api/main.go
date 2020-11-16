package main

import (
	"log"

	routes "github.com/EvgenyiK/aircrafts/api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router:= gin.Default()
	routes.Routes(router)
	log.Fatal(router.Run(":8080"))
}