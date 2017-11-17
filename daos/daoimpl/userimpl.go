package daoimpl

import (
	"OnlineTestGo/models"
	"OnlineTestGo/tos"
	"OnlineTestGo/utility"
	"fmt"
	"log"
)

type UserImpl struct{}

func (dao UserImpl) SaveNewUser(user models.User) (tos.Userto, error) {

	utility.GetLogger()

	log.Println("entering in SaveNewUser() function")
	var newuser tos.Userto

	log.Println("executing query and storing registration details")
	query := "insert into registration(fname, lname, phone, email,password,usertype) values(?,?,?,?,?,?)"
	db := connection()
	defer db.Close()
	stmt, err := db.Prepare(query)
	log.Println(err)

	defer stmt.Close()

	res, err := stmt.Exec(user.Fname, user.Lname, user.Phone, user.Email, user.Password, user.UserType)
	//var usertos tos.User
	newuser.Message = " registration successful"
	if err != nil {
		log.Panic("Exec err:", err.Error())
	}

	id, err := res.LastInsertId()
	newuser.ID = id
	//newuser = append(newuser, usertos)
	fmt.Printf("%T", id)
	if err != nil {
		log.Println("Exec err:", err.Error())
	}

	return newuser, err
}

func (dao UserImpl) CheckUser(user models.User) (tos.Userto, error) {
	utility.GetLogger()
	log.Println("entering in userDao.CheckUser()...")
	//var existuser []tos.User
	phone := user.Phone
	log.Println("executing query and checking user exists")
	query := "select id from registration where phone = ?"
	db := connection()
	defer db.Close()

	rows, err := db.Query(query, phone)
	if err != nil {
		log.Fatal(err)
		defer rows.Close()
	}
	var usertos tos.Userto
	usertos.Message = "you are an existing user  check the existing credentials"
	for rows.Next() {

		err := rows.Scan(&usertos.ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(usertos)
		//existuser = append(existuser, usertos)
	}

	//var usertos tos.User
	return usertos, err
}

func (dao UserImpl) AuthenticateUser(user models.User) models.User {

	utility.GetLogger()

	log.Println("entering in AuthenticateUser()")

	log.Println("Executing query and authenticating user exist")
	query := "select id, fname, lname, phone, usertype from registration where email=? and password = ?"
	db := connection()
	defer db.Close()

	rows, err := db.Query(query, user.Email, user.Password)

	if err != nil {
		log.Fatal(err)
		defer rows.Close()
	}

	var newuser models.User
	for rows.Next() {

		err := rows.Scan(&newuser.ID, &newuser.Fname, &newuser.Lname, &newuser.Phone, &newuser.UserType)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(newuser)
	}

	log.Println("Response user Obj : ", newuser)

	return newuser
}
