package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	utility.GetLogger()
	log.Println("entering in webservice.RegisterUser")
	var user models.User
	c.BindJSON(&user)

	log.Println("calling manager.Register function")
	insertedid := manager.Register(user)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": insertedid,
	})
}
