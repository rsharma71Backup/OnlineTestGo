package webservice

import (
	"OnlineTestGo/models"
	"OnlineTestGo/tos"

	"github.com/gin-gonic/gin"
)

func Mocklogin(c *gin.Context) {

	var user models.User
	var token tos.Tokento
	c.BindJSON(&user)

	responsemessage := "failure"
	if user.Email == "test@t.com" {
		token = tos.Tokento{Uid: 24, Fname: "umashankar", Token: "dadasd545a4d", UserType: "user"}
		responsemessage = "success"
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  responsemessage,
		"message": token,
	})
}
