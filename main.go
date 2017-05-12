package main

import (
	"OnlineTestGo/webservice"

	"OnlineTestGo/utility"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/goinggo/tracelog"
)

func main() {
	utility.GetLogger()
	tracelog.Start(tracelog.LevelTrace)
	tracelog.Trace("main", "main", " Trace")
	tracelog.Info("main", "main", " Info")
	tracelog.Warning("main", "main", " Warn")
	tracelog.Errorf(fmt.Errorf("Exception At..."), "main", "main", " Error")

	router := gin.Default()
	router.GET("/testService", webservice.TestService)
	router.POST("/registerUser", webservice.RegisterUser)
	router.POST("/login", webservice.Login)

	router.GET("/questions", webservice.QuestionList)
	router.POST("/userAnswer", webservice.AnswerList)

	router.GET("/admin", webservice.Admin)
	router.POST("/addquestions", webservice.AddQuestions)

	router.GET("/logout", webservice.Logout)
	router.GET("/mockadmin", webservice.MockAdmin)

	router.Run(GetPort())

}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "7082"
		port = "8084"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
