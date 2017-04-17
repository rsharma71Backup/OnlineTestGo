package webservice

import (
	"OnlineTestGo/models"

	"github.com/gin-gonic/gin"
)

func Mocklogout(c *gin.Context) {

	var token models.Token
	c.BindJSON(&token)
	message := "sucess"
	if token.Token == "ghctfdtf" {
		message = "sucess"
	}

	message = "sucess"

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  message,
		"message": "logout",
	})

}
