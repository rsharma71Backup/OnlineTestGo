package webservice

import (
	"OnlineTestGo/tos"

	"github.com/gin-gonic/gin"
)

func MockAdmin(c *gin.Context) {

	var dlist tos.Admin1

	dlist = tos.Admin1{Uid: 1, Fname: "Sweta", Lname: "Vahia", Java: 20, Fundamental: 30, Language: 25}

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(200, gin.H{
		"status":  "successs",
		"message": dlist,
	})
}
