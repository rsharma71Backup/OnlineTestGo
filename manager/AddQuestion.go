package manager

import (
	"OnlineTestGo/daos/daoimpl"

	"OnlineTestGo/models"

	"log"

	"OnlineTestGo/utility"
)

func AddQuestion(question models.Question) int64 {
	utility.GetLogger()
	log.Println("entering into manager.AddQuestion")

	questionDao := daoimpl.QuestionImpl{}
	log.Println("calling questionDao.AddQuestion()")

	insertedid, err := questionDao.AddQuestion(question)
	if err != nil {
		log.Println("error occured", err)
	}
	log.Println(insertedid)
	return insertedid
}
