package main

import (
	"bytes"
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

	type Question struct {
		Id       int    `json:"id"`
		Question string `json:"question"`
		Option1  string `json:"option1"`
		Option2  string `json:"option2"`
		Option3  string `json:"option3"`
		Option4  string `json:"option4"`
		Option5  string `json:"option5"`
		Answer   string `json:"answer"`
		Type     string `json:"type"`
	}
	router := gin.Default()
	//Get one question
	router.GET("/question/:id", func(c *gin.Context) {
		var (
			question Question
			result   gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select * from questions where id=?;", id)
		err = row.Scan(&question.Id, &question.Question, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": question,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	//Get all questions
	router.GET("/questions", func(c *gin.Context) {
		var questions []Question
		rows, err := db.Query("select * from questions")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			var question Question
			err = rows.Scan(&question.Id, &question.Question, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
			questions = append(questions, Question{question.Id, question.Question, question.Option1, question.Option2, question.Option3, question.Option4, question.Option5, question.Answer, question.Type})

			if err != nil {
				fmt.Print(err.Error())
			}
		}

		defer rows.Close()

		c.JSON(http.StatusOK, gin.H{
			"result": questions,
			"count":  len(questions),
		})
	})

	package main

import (
	"bytes"
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

	type Question struct {
		Id       int    `json:"id"`
		Question string `json:"question"`
		Option1  string `json:"option1"`
		Option2  string `json:"option2"`
		Option3  string `json:"option3"`
		Option4  string `json:"option4"`
		Option5  string `json:"option5"`
		Answer   string `json:"answer"`
		Type     string `json:"type"`
	}
	router := gin.Default()
	//Get one question
	router.GET("/question/:id", func(c *gin.Context) {
		var (
			question Question
			result   gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select * from questions where id=?;", id)
		err = row.Scan(&question.Id, &question.Question, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
		if err != nil {
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": question,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})
	//Get all questions
	router.GET("/questions", func(c *gin.Context) {
		var questions []Question
		rows, err := db.Query("select * from questions")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			var question Question
			err = rows.Scan(&question.Id, &question.Question, &question.Option1, &question.Option2, &question.Option3, &question.Option4, &question.Option5, &question.Answer, &question.Type)
			questions = append(questions, Question{question.Id, question.Question, question.Option1, question.Option2, question.Option3, question.Option4, question.Option5, question.Answer, question.Type})

			if err != nil {
				fmt.Print(err.Error())
			}
		}

		defer rows.Close()

		c.JSON(http.StatusOK, gin.H{
			"result": questions,
			"count":  len(questions),
		})
	})

	//Post New Question
	router.POST("/question", func(c *gin.Context) {
		var question Question
		c.BindJSON(&question)
		stmt, err := db.Prepare("insert into questions(question, option1, option2, option3, option4, option5, answer, type values(?,?,?,?,?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(question.Question, question.Option1, question.Option2, question.Option3, question.Option4, question.Option5, question.Answer, question.Type)

		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Question created Successfully"),
		})
	})

	// PUT - update a question
	router.PUT("/question", func(c *gin.Context) {
		var buffer bytes.Buffer
		Id := c.Query("id")
		Question := c.PostForm("question")
		Option1 := c.PostForm("option1")
		Option2 := c.PostForm("option2")
		Option3 := c.PostForm("option3")
		Option4 := c.PostForm("option4")
		Option5 := c.PostForm("option5")
		stmt, err := db.Prepare("update questions set question=?, option1= ?, option2=?, option3=?, option4=?, option5=? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(Question, Option1, Option2, Option3, Option4, Option5, Id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(Question)
		buffer.WriteString("\n")
		buffer.WriteString("Option1 ")
		buffer.WriteString(Option1)
		buffer.WriteString("\nOption2 ")
		buffer.WriteString(Option2)
		buffer.WriteString("\nOption3 ")
		buffer.WriteString(Option3)
		buffer.WriteString("\nOption4 ")
		buffer.WriteString(Option4)
		buffer.WriteString("\nOption5 ")
		buffer.WriteString(Option5)
		defer stmt.Close()
		que := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", que),
		})
	})

	// Delete resources
	router.DELETE("/question", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from questions where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted question: %s", id),
		})
	})
	router.Run(":9092")
}

	// PUT - update a question
	router.PUT("/question", func(c *gin.Context) {
		var buffer bytes.Buffer
		Id := c.Query("id")
		Question := c.PostForm("question")
		Option1 := c.PostForm("option1")
		Option2 := c.PostForm("option2")
		Option3 := c.PostForm("option3")
		Option4 := c.PostForm("option4")
		Option5 := c.PostForm("option5")
		stmt, err := db.Prepare("update questions set question=?, option1= ?, option2=?, option3=?, option4=?, option5=? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(Question, Option1, Option2, Option3, Option4, Option5, Id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(Question)
		buffer.WriteString("\n")
		buffer.WriteString("Option1 ")
		buffer.WriteString(Option1)
		buffer.WriteString("\nOption2 ")
		buffer.WriteString(Option2)
		buffer.WriteString("\nOption3 ")
		buffer.WriteString(Option3)
		buffer.WriteString("\nOption4 ")
		buffer.WriteString(Option4)
		buffer.WriteString("\nOption5 ")
		buffer.WriteString(Option5)
		defer stmt.Close()
		que := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", que),
		})
	})

	// Delete resources
	router.DELETE("/question", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from questions where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(id)
		if err != nil {
			fmt.Print(err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully deleted question: %s", id),
		})
	})
	router.Run(":9092")
}
