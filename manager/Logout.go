package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/utility"

	"log"
)

func DeleteToken(delt string) string {
	utility.GetLogger()
	log.Println("entering manager.DeleteToken()")

	logoutdao := daoimpl.LogoutImpl{}
	log.Println("calling logoutdao.DeleteToken()")
	a := logoutdao.DeleteToken(delt)

	return a
}
