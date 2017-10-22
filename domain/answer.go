package domain

import "time"

type Answer struct {
	Id         int
	UserId     string
	QuestionId string
	Answer     string
	Timestamp  time.Time
}
