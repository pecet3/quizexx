package data

import (
	"github.com/pecet3/quizex/data/dtos"
)

func (f FunFact) ToDTO(q *Queries) *dtos.FunFact {
	return &dtos.FunFact{
		Topic:   f.Topic,
		Content: f.Content,
	}
}
