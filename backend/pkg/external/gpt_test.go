package external

import (
	"context"
	"fmt"
	"testing"
)

func TestFetchQuestionSet(t *testing.T) {
	category := "cars"
	maxRounds := "4"
	difficulty := "medium"
	lang := "english"

	service := ExternalService{}

	questionSet, err := service.FetchQuestionSet(context.Background(), category, maxRounds, difficulty, lang)

	if err != nil {
		t.Errorf("error: %v", err)
	}

	if questionSet == "" {
		t.Errorf("empty answer")
	}

	fmt.Println("Question set:")
	fmt.Println(questionSet)
}
