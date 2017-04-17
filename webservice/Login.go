package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	utility.GetLogger()
	log.Println("entering webservice.Login()")

	var user models.User

	c.BindJSON(&user)
	log.Println("calling manager.Login()")
	log.Println(user)
	message := manager.Login(user)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": message,
	})
}
