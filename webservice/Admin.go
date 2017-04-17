package webservice

import (
	"OnlineTestGo/manager"
	"OnlineTestGo/utility"
	"log"

	"github.com/gin-gonic/gin"
)

func Admin(c *gin.Context) {
	utility.GetLogger()
	log.Println("entering webservice.Admin()")
	log.Println("calling manager.FetchData()")

	dlist := manager.FetchData()
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status": "sucess",
		"score":  dlist,
	})

}
