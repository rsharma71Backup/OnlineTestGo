package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func FetchQuestions(c *gin.Context) {

	utility.GetLogger()
	log.Println("entering into manager.()")
	log.Println("calling manager.AdminAddQuestion()")
	//var question models.Question
	//err := c.BindJSON(&question)
	//log.Println("Error", err)
	//if err != nil {
	//log.Println("Error", err)
	//}

	log.Println("calling manager.FetchQuestion()")
	Qlist := manager.AdminFetchQuestion()
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "success",
		"message": Qlist,
	})
}
