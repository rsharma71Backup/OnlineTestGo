package webservice

import (

	//	"OnlineTestGo/manager"

	"github.com/gin-gonic/gin"
	//	"OnlineTestGo/models"
)

func TestService(c *gin.Context) {

	c.JSON(200, gin.H{
		"status":  "success",
		"message": "your webserivce is reachable",
	})

}
