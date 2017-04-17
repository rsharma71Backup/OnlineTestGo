package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"log"
)

var userDao interfaces.UserDao

func Register(user models.User) tos.Userto {

	utility.GetLogger()
	log.Println("entering in  manager.Register() function")

	log.Println("calling register manager")
	userDao := daoimpl.UserImpl{}
	log.Println("calling userDao.CheckUser function")
	existuser, err := userDao.CheckUser(user)

	id := existuser.ID
	{
		if id != 0 {
			return existuser
			log.Println(id)
		}
	}

	//user.UserType = "user"
	log.Println("calling userDao.SaveNewUser() function")
	newuser, err := userDao.SaveNewUser(user)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(newuser)

	return newuser
}
