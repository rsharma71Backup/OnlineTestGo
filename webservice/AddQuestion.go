package webservice

import (
	"OnlineTestGo/filter"
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func AddQuestions(c *gin.Context) {
	utility.GetLogger()

	log.Println("entering in webservice.AddQuestions()")
	token := c.Query("Authorization")
	log.Println("calling filter.AuntheticateToken()")
	bool := filter.AuthenticateToken(token)
	log.Println(bool)

	if bool == true {
		var question models.Question

		c.BindJSON(&question)

		log.Println("calling manager.AddQuestion()")

		insertedid := manager.AddQuestion(question)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(200, gin.H{
			"status":  "success",
			"message": insertedid,
		})

	} else {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		c.JSON(401, gin.H{
			"status":  "failure",
			"message": "you dont have permission to acces",
		})

	}
}
