package tos


type Question struct {
	ID            int      `json:"id"`
	Qno           int      `json:"qno"`
	Question      string   `json:"question"`
	Options       []string `json:"choices"`
	CorrectAnswer string   `json:"correctAnswer"`
	Type          string   `json:"type"`
}
