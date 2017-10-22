package domain_test

import (
	"testing"

	"github.com/VagabondDataNinjas/gizlinebot/domain"
)

var ids = []string{"Q_1", "Q_2", "Q_3"}
var questions = []string{
	"what is your name",
	"where do you live",
	"what is your hobby",
}

func makeQuestions() *domain.Questions {
	qs := domain.NewQuestions()
	for i, _ := range ids {
		qs.Add(ids[i], questions[i])
	}

	return qs
}

func TestAt(t *testing.T) {
	qs := makeQuestions()
	for j, _ := range ids {
		a, err := qs.At(j)
		if err != nil {
			t.Fatal(err)
		}

		if a.Id != ids[j] {
			t.Fatalf("Expected %s, got %s", ids[j], a.Id)
		}

		if a.Text != questions[j] {
			t.Fatalf("Expected %s, got %s", questions[j], a.Text)
		}
	}

	nonExistent, err := qs.At(len(ids))
	if err == nil {
		t.Fatal("Expected err but got nil")
	}

	if nonExistent != nil {
		t.Fatal("Expected nil but for a value At()")
	}
}

func TestNext(t *testing.T) {
	qs := makeQuestions()

	next, err := qs.Next(ids[0])
	if err != nil {
		t.Fatal(err)
	}

	if next.Id != ids[1] || next.Text != questions[1] {
		t.Fatalf("Unexpected values for next: %+v", next)
	}

	_, err = qs.Next(ids[len(ids)-1])
	if err == nil {
		t.Fatal("Expected err but got nil")
	}
	if err != domain.ErrQuestionsNoNext {
		t.Fatalf("Expected ErrQuestionsNoNext but got %T", err)
	}
}
