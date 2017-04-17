package webservice

import (
	"OnlineTestGo/manager"

	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	utility.GetLogger()
	log.Println("entering webservice.Logout()")
	token := c.Query("Authorization")
	log.Println("calling manager.DeleteTokenLogin()")
	message := manager.DeleteToken(token)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "sucess",
		"message": message,
	})

}
