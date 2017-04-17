
package main

import (
	//"bytes"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	db, err := sql.Open("mysql", "Rakesh:Root12345$@tcp(rpqb.centralindia.cloudapp.azure.com:3306)/interview_test")
	if err != nil {
		fmt.Print(err.Error())
	}

	defer db.Close()

// make sure connection is available

	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Answer struct {
		Id   		int	`json:"id"`
		Userid     	int	`json:"userid"`
		Questionid      int	`json:"questionid"`
		Answer      	string	`json:"answer"`
		Correctness 	string	`json:"correctness"`
	}
	router := gin.Default()

	// GET a answer of one person
	router.GET("/answer/:id", func(c *gin.Context) {
		var (
			answer Answer
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select * from answers where id=?;", id)
		err = row.Scan(&answer.Id, &answer.Userid, &answer.Questionid, &answer.Answer, &answer.Correctness)
		if err != nil {
			//if no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": answer,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	//get all the answers
	router.GET("/answers", func(c *gin.Context) {
		var (
			answer  Answer
			answers []Answer
		)
		rows, err := db.Query("select * from answers;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			err = rows.Scan(&answer.Id, &answer.Userid, &answer.Questionid, &answer.Answer, &answer.Correctness)
			answers = append(answers, Answer{answer.Id, answer.Userid, answer.Questionid, answer.Answer, answer.Correctness}) 
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": answers,
			"count":  len(answers),
		})
	})

	//Inserting new answers to database

	router.POST("/answer", func(c *gin.Context) {
		var answer Answer 		
		c.BindJSON(&answer)
		stmt, err := db.Prepare("insert into answers( userid, questionid, answer, correctness) values(?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(answer.Userid,answer.Questionid,answer.Answer,answer.Correctness)
	     c.JSON(http.StatusOK, gin.H{
		         "message": fmt.Sprintf("Answer successfully created"),
		})
	})
	// Delete resources
	router.DELETE("/answer", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from answers where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted user: %s", id),
		})
	})

	router.Run(":9091")
}

