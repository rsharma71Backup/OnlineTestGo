package daoimpl

import (
	"OnlineTestGo/utility"

	"log"
)

type LogoutImpl struct{}

func (dao LogoutImpl) DeleteToken(token string) string {
	utility.GetLogger()
	log.Println("exntering in DeleteToken() function")
	deltoken := "notfoud"
	log.Println("executing query and deleting token from database ")
	query := "DELETE FROM token WHERE token=?"

	db, conn := connectaws()
	defer db.Close()
	defer conn.Close()
	stmt, err := db.Prepare(query)
	if err != nil {
		return deltoken
	}

	defer stmt.Close()
	log.Println(token)
	res, err1 := stmt.Exec(token)

	if err1 != nil {
		log.Panic("Exec err:", err1.Error())
	}
	q, err2 := res.RowsAffected()
	if err2 != nil {
		log.Panic("Exec err:", err2.Error())
	}
	log.Println(q)
	if q > 0 {
		deltoken = "deleted"
	}

	return deltoken

}
