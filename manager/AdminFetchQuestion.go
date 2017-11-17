package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"log"
)

func AdminFetchQuestion() []tos.Question {

	utility.GetLogger()
	log.Println("entering into manager.FetchAllQuestions")

	log.Println("calling Question manager")

	questionDao1 := daoimpl.QuestionImpl{}

	log.Println("calling FetchAllQuestions()")

	questionlist := questionDao1.FetchAllQuestions()

	return questionlist

}
