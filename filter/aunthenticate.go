package filter

import (
	"OnlineTestGo/daos/daoimpl"
	"log"

	"OnlineTestGo/utility"
	"time"
)

func AuthenticateToken(token string) bool {
	utility.GetLogger()
	log.Println("entering into AuthenticateToken()")
	dummytime := "Dec 29, 2014 at 6:00am (SGT)"
	layOut := "Jan 2, 2006 at 3:04am (MST)"
	timeStamp, err := time.Parse(layOut, dummytime)

	if err != nil {
		log.Println(err)

	}

	hr, min, sec := timeStamp.Clock()
	log.Println(hr, min, sec)

	var tokenEncodeString string

	if len(token) > 0 {
		tokenEncodeString = token
	}

	tokenDao := daoimpl.TokenImpl{}

	log.Println("calling  daoimpl tokenDao.AunthenticateToken()")
	tokenfromdb, lastaccesstime := tokenDao.AunthenticateToken(tokenEncodeString)

	if len(tokenfromdb) == 0 {
		return false

	}
	currenttime := time.Now().Format("2006-01-02 15:04:05")
	log.Println(currenttime)
	const layOut1 = "2006-01-02 15:04:05"
	currentaccesstime, err := time.Parse(layOut1, currenttime)
	if err != nil {
		log.Println(err)

	}

	duration := currentaccesstime.Sub(lastaccesstime)

	if tokenEncodeString == tokenfromdb {

		log.Println("calling ModifyLastAccessTime()")
		tokenDao.ModifyLastAccessTime(currentaccesstime, tokenEncodeString)

	}
	if duration.Hours() > float64(hr) {

		log.Println("calling DeleteToken() function")
		bool := tokenDao.DeleteToken(tokenEncodeString)
		log.Println(bool)

	}
	return true
}
