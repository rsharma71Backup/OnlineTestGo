package webservice

import (
	"fmt"

	"log"

	"OnlineTestGo/manager"

	"OnlineTestGo/utility"

	"github.com/gin-gonic/gin"

	"OnlineTestGo/models"
)

func AnswerList(c *gin.Context) {
	utility.GetLogger()

	log.Println("entering webservice.AnswerLIst()")
	var answer []models.Answer
	c.BindJSON(&answer)
	log.Println(answer)
	log.Println("calling manager.CalculateScore()")
	score := manager.CalculateScore(answer)

	fmt.Println("answerlist", score)

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{

		"status": "success",

		"score": score,
	})
}
