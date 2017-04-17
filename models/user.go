package models

type User struct {
	ID    int    `json:"id"`
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	Password string `json:"password"`
	UserType string `json:"usertype"`
}
