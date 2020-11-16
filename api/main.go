package main

import (
	"github.com/rs/cors/wrapper/gin"
	routes "aircrafts/api/routes"
)

func main() {
	router:= gin.Default()
	routes
}