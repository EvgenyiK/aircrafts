package main

import (
	"log"

	"github.com/EvgenyiK/aircrafts/controllers"
)

func main() {
	r := controllers.Find()
	log.Fatal(r.Run(":8080"))
}
