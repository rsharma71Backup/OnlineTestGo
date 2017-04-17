package manager

import (
	"OnlineTestGo/daos/daoimpl"
	"OnlineTestGo/daos/interfaces"
	"OnlineTestGo/models"
	"OnlineTestGo/utility"
	"log"
	"math/big"
	"time"
)

var questionDao interfaces.QuestionDao

func CalculateScore(answerList []models.Answer) int {
	start := time.Now()

	utility.GetLogger()
	r := new(big.Int)
	log.Println(r.Binomial(1000, 10))
	log.Println("entering manager.CalculateScore()")

	log.Println("calling Answer manager")

	questionDao := daoimpl.QuestionImpl{}
	answerDao := daoimpl.AnswerImpl{}

	score := 0

	log.Println("answerList", answerList)
	for _, answer := range answerList {

		log.Println("calling questionDao.GetAnswerById() and Fetching the right answer...")
		start1 := time.Now()
		r1 := new(big.Int)
		log.Println(r1.Binomial(1000, 10))
		correctAnswer := questionDao.GetAnswerById(answer.Q_type)
		elapsed1 := time.Since(start1)
		log.Println("Binomial took %s", elapsed1)

		log.Println("actual vs correct", answer.Selected, correctAnswer)
		if answer.Selected == correctAnswer {

			score++
			start2 := time.Now()
			r2 := new(big.Int)
			log.Println(r2.Binomial(1000, 10))
			answerDao.SaveAnswer(answer)
			elapsed2 := time.Since(start2)
			log.Println("Binomial took %s", elapsed2)
		}

	}
	elapsed := time.Since(start)
	log.Println("Binomial took %s", elapsed)
	return score

}
