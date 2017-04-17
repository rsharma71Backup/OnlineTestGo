package webservice

import (
	"OnlineTestGo/filter"
	"OnlineTestGo/manager"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func QuestionList(c *gin.Context) {

	utility.GetLogger()
	log.Println("entering into manager.QuestionList()")
	testtype := c.Query("testtype")
	token := c.Query("Authorization")
	log.Println(testtype)
	log.Println(token)
	log.Println("calling filter.AuntheticateToken()")
	bool := filter.AuthenticateToken(token)

	if bool == true {
		log.Println("calling manager.FetchQuestion()")
		Qlist := manager.FetchQuestion(testtype)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

		c.JSON(200, gin.H{
			"status":  "success",
			"message": Qlist,
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
