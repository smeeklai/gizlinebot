package domain

import (
	"fmt"

	"github.com/pkg/errors"
)

var ErrQuestionsNoNext = errors.New("No next question")

// stores a list of Question objects
// not a map because we need to keep the order of the questions
// maps in Go do not keep order
type Questions struct {
	text []string
	ids  []string
}

func NewQuestions() *Questions {
	return &Questions{}
}

func (qs *Questions) Add(id, question string) *Questions {
	qs.text = append(qs.text, question)
	qs.ids = append(qs.ids, id)
	return qs
}

func (qs *Questions) At(index int) (q *Question, err error) {
	if index >= len(qs.ids) {
		return q, errors.New(fmt.Sprintf("Index [%d] out of range", index))
	}
	return &Question{
		Id:   qs.ids[index],
		Text: qs.text[index],
	}, nil
}

func (qs *Questions) Next(qid string) (q *Question, err error) {
	returnNext := false
	for index, id := range qs.ids {
		if returnNext {
			return qs.At(index)
		}
		if id == qid {
			returnNext = true
			continue
		}
	}
	return q, ErrQuestionsNoNext
}
