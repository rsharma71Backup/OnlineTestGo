package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"log"
)

func FetchQuestion(testtype string) []tos.Question {

	utility.GetLogger()
	log.Println("entering into manager.FetchQuestion")

	log.Println("calling Question manager")

	questionDao1 := daoimpl.QuestionImpl{}

	log.Println("calling FetchQuestionsByType()")

	questionlist := questionDao1.FetchQuestionsByType(testtype)

	return questionlist

}
