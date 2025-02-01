package data

import (
	"context"
	"math"

	"github.com/pecet3/quizex/data/dtos"
	"github.com/pecet3/quizex/pkg/logger"
)

func (u User) ToDto(d *Queries) *dtos.User {
	ug, err := d.GetGameUserByUserID(context.Background(), u.ID)
	if err != nil {
		logger.Error(err)
		return nil
	}

	return &dtos.User{
		Name:      u.Name,
		UUID:      u.Uuid,
		Email:     u.Email.String,
		CreatedAt: u.CreatedAt.Time,
		ImageUrl:  u.ImageUrl,
		Progress:  math.Round(ug.Progress*100) / 100,
		Level:     int(ug.Level),
		Exp:       math.Round(ug.Exp*100) / 100,
	}
}
