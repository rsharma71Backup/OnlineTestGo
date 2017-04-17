package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"Rakesh:Root12345$@tcp(rpqb.centralindia.cloudapp.azure.com:3306)/interview_test")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	type Person struct {
		Id string `json:"id"`
		Fname string `json:"fname"`
		Lname string `json:"lname"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}
	router := gin.Default()
	// GET a person detail
	router.GET("/person/:id", func(c *gin.Context) {
		var (
			person Person
			result gin.H
		)
		id := c.Param("id")
		row := db.QueryRow("select * from registration where id =?;", id)
		err = row.Scan(&person.Id, &person.Fname, &person.Lname, &person.Phone, &person.Email)
		if err != nil {
			// If no results send null
			result = gin.H{
				"result": nil,
				"count":  0,
			}
		} else {
			result = gin.H{
				"result": person,
				"count":  1,
			}
		}
		c.JSON(http.StatusOK, result)
	})

	// GET all persons
	router.GET("/persons", func(c *gin.Context) {
		var persons []Person
		rows, err := db.Query("select * from registration;")
		if err != nil {
			fmt.Print(err.Error())
		}
		for rows.Next() {
			var person Person
			err = rows.Scan(&person.Id, &person.Fname, &person.Lname, &person.Phone, &person.Email)
			//persons = append(persons, Person{id: person.id, fname: person.Fname, lname: person.lname, phone: person.phone, email: person.email})
			persons = append(persons, Person{person.Id, person.Fname, person.Lname, person.Phone, person.Email})
			if err != nil {
				fmt.Print(err.Error())
			}
		}
		defer rows.Close()
		c.JSON(http.StatusOK, gin.H{
			"result": persons,
			"count":  len(persons),
		})
	})

	// POST new person details
	router.POST("/person", func(c *gin.Context) {
		var person Person
		c.BindJSON(&person)
		stmt, err := db.Prepare("insert into registration(fname, lname, phone, email) values(?,?,?,?);")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(person.Fname,person.Lname,person.Phone,person.Email) 

 
 		if err != nil { 
			fmt.Print(err.Error()) 
 		} 

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("successfully created %v", person.Fname),
		})
	})

	// PUT - update a person details
	router.PUT("/person", func(c *gin.Context) {
		var buffer bytes.Buffer
		Id := c.Query("id")
		Fname := c.PostForm("fname")
		Lname := c.PostForm("lname")
		// phone := c.PostForm("phone")
		// email := c.PostForm("email")
		stmt, err := db.Prepare("update registration set fname= ?, lname= ? where id= ?;")
		if err != nil {
			fmt.Print(err.Error())
		}
		_, err = stmt.Exec(Fname, Lname, Id)
		if err != nil {
			fmt.Print(err.Error())
		}

		// Fastest way to append strings
		buffer.WriteString(Fname)
		buffer.WriteString(" ")
		buffer.WriteString(Lname)
		defer stmt.Close()
		name := buffer.String()
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	})

	// Delete resources
	router.DELETE("/person", func(c *gin.Context) {
		id := c.Query("id")
		stmt, err := db.Prepare("delete from registration where id= ?;")
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
	router.Run(":9090")
}
