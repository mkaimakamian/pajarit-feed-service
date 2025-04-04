package createpost

import (
	"pajarit-feed-service/domain"
	"time"
)

type CreatePostCmd struct {
	AuthorId string
	Content  string
}

type CreatePostResponse struct {
	Id        string    `json:"id"`
	AuthorId  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCreatePostResponse(post *domain.Post) *CreatePostResponse {
	// Realizo un casteo directo por practicidad; en muhcos casos
	// esto no sería posible (y hasta sería considerado una mala práctica)
	response := CreatePostResponse(*post)
	return &response
}
