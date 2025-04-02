package createpost

import (
	"pajarit-feed-service/domain"
	"time"
)

type CreatePostCmd struct {
	AuthorId string
	Message  string
}

type CreatePostResponse struct {
	Id           string
	AuthorId     string
	Message      string
	CreationDate time.Time
}

func NewCreatePostResponse(post *domain.Post) *CreatePostResponse {
	// Realizo un casteo directo por practicidad; en muhcos casos
	// esto no sería posible (y hasta sería considerado una mala práctica)
	response := CreatePostResponse(*post)
	return &response
}
