package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func AdminAddQuestions(c *gin.Context) {

	utility.GetLogger()
	log.Println("entering into manager.AdminAddQuestion()")
	log.Println("calling manager.AdminAddQuestion()")
	var question models.Question

	c.BindJSON(&question)
	log.Println("calling manager.AdminAddQuestion()")

	insertedid := manager.AddQuestion(question)
	log.Println(insertedid)
	log.Println("calling manager.FetchQuestion()")
	Qlist := manager.AdminFetchQuestion()
	c.Header("Access-Control-Allow-Origin", "*")

	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	c.JSON(200, gin.H{
		"status":  "success",
		"message": Qlist,
	})
}
