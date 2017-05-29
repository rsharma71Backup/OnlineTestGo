package daoimpl

import (
	"OnlineTestGo/models"
	"log"
)

type AnswerImpl struct{}

func (dao AnswerImpl) SaveAnswer(answer models.Answer) (int64, error) {

	query := "INSERT into Answers(uid,q_type,score) values (?,?,?)"

	db := connection()
	defer db.Close()
	//defer conn.Close()

	stmt, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(answer.Uid, answer.Uid, 1)

	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return id, err

}
