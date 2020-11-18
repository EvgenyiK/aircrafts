package controllers

import (
	"net/http"
	"github.com/EvgenyiK/aircrafts/models"
	"github.com/gin-gonic/gin"
)

var air models.Aircraft

func FindAircrafts(c *gin.Context){
	var aircrafts []models.Aircraft
	models.DB.Table("aircrafts").First(&aircrafts)
	models.DB.Find(&aircrafts)
	//c.JSON(http.StatusOK, gin.H{"data": aircrafts})
	for _, v := range aircrafts {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"name": v.Model,
			"range": v.Range,
			"code" : v.Code,
		})
	}
}