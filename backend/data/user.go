package data

import "github.com/pecet3/quizex/data/dtos"

func (u User) ToDto(d *Queries) *dtos.User {
	return &dtos.User{
		Name:     u.Name,
		UUID:     u.Uuid,
		Email:    u.Email.String,
		ImageUrl: u.ImageUrl,
	}
}
