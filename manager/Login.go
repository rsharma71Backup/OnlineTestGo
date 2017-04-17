package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func Login(user models.User) tos.Tokento {
	utility.GetLogger()
	log.Println("entering manager.login")
	var tokenObj tos.Tokento
	var tokenModel models.Token

	userDao := daoimpl.UserImpl{}
	tokenDao := daoimpl.LoginImpl{}
	log.Println("calling userDao.AuthenticateUser()")
	userObj := userDao.AuthenticateUser(user)

	if userObj.ID != 0 {
		log.Println("generates token if user exist by calling GenerateToken()  ")
		token := GenerateToken()

		tokenModel.LastAccessTime = time.Now()
		tokenModel.Token = token
		tokenModel.Uid = userObj.ID
		log.Println("calling tokenDao.DeleteDuplicateUid()")
		tokenDao.DeleteDuplicateUid(tokenModel)
		log.Println("calling tokenDao.SaveToken()")

		id, err := tokenDao.SaveToken(tokenModel)
		if err != nil {
			log.Println(err)
		}

		if id != 0 {
			tokenObj.Token = token
			tokenObj.Fname = userObj.Fname
			tokenObj.Uid = userObj.ID
			tokenObj.UserType = userObj.UserType
		}

		log.Println(tokenObj)

	}

	return tokenObj

}

func GenerateToken() string {
	str := RandStringRunes(20)
	{
		fmt.Println(str)

	}
	return str
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	token := make([]rune, n)
	for i := range token {
		token[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(token)
}
