package survey

import (
	"github.com/VagabondDataNinjas/gizlinebot/domain"
	"github.com/VagabondDataNinjas/gizlinebot/storage"
)

type Survey struct {
	Storage   storage.Storage
	Questions *domain.Questions
}

func NewSurvey(storage storage.Storage, questions *domain.Questions) (survey *Survey) {
	return &Survey{
		Storage:   storage,
		Questions: questions,
	}
}

func (s *Survey) GetNextQuestion(userId string) (question *domain.Question, err error) {
	has, err := s.Storage.UserHasAnswers(userId)
	if err != nil {
		return question, err
	}
	if !has {
		return s.Questions.At(0)
	}

	answer, err := s.Storage.GetUserLastAnswer(userId)
	if err != nil {
		return question, err
	}

	answer.QuestionId
}
