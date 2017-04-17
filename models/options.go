package models

//indicates options table
type Options struct {
	ID      int    `json:"id"`
	Qid     int    `json:"qid"`
	Choices string `json:"choices"`
	Answers string `json:"answers"`
}
