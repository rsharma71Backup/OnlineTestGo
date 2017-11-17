package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"log"
)

func FetchData() []tos.Admin {
	utility.GetLogger()
	log.Println("entering manager.FetchData()")
	adminDao := daoimpl.AdminImpl{}
	log.Println("calling adminDao.FetchData()")
	return adminDao.FetchData()

}
