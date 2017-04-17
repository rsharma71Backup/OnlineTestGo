package models

//Answer indicates answer table
type Answer struct {
	Uid      int64  `json:"uid"`
	Q_type   int64  `json:"q_type"`
	Selected string `json:"selected"`
}