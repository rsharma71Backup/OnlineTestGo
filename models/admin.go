package models

type Admin struct {
	Uid      int    `json:"uid"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Testtype string `json:"testtype"`
	Score    int    `json:"score"`
}
