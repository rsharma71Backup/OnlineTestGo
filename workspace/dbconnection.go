package main

import (
	//"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql",
		"Rakesh:Root12345$@tcp(rpqb.centralindia.cloudapp.azure.com:3306)/interview_test")
	if err != nil {
		log.Fatal(err)
	}
	insertlist(db)
	getlist(db)
	defer db.Close()

}
func insertlist(db *sql.DB) {
	rows, err := db.Query("insert into registration(id,fname,lname,phone,email) values(129,'Rakesh','kumar','9681817926','r@rapidqube.com')")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	err = rows.Err()
}
func getlist(db *sql.DB) {
	var (
		id    int
		fname string
		lname string
		phone int
		email string
	)
	rows, err := db.Query("select * from registration")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &fname, &lname, &phone, &email)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, fname, lname, phone, email)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
