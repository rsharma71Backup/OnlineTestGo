package interfaces

import (
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
)

type QuestionDao interface {
	FetchQuestionsByType(testtype string) ([]tos.Question, error)

	AddQuestion(question models.Question) (int64, error)
}
