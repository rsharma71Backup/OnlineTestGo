package tos

type Admin struct {
	Uid   int      `json:"uid"`
	Fname string   `json:"fname"`
	Lname string   `json:"lname"`
	Type  []string `json:"type"`
	Score []int    `json:"score"`
}
