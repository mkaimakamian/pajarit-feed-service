package gettimeline

import (
	"pajarit-feed-service/domain"
	"time"
)

type GetTimelineCmd struct {
	UserId string
	Offset int
	Limit  int
	Size   int
}

type TimelineResponse struct {
	UserId string
	Feed   []PostResponse

	Offset int
	Limit  int
	Size   int
}

// Si bien podría haber referenciado a a createpost.CreatePostResponse, previa
// reubicación en un paquete común, prioricé el aislamiento por resultar más
// práctico para el challenge.
type PostResponse struct {
	Id           string
	AuthorId     string
	Message      string
	CreationDate time.Time
}

func NewGetTimelineResponse(timeline *domain.Timeline, userId string) *TimelineResponse {

	response := TimelineResponse{UserId: userId}

	for _, post := range timeline.Posts {
		post := PostResponse(post)
		response.Feed = append(response.Feed, post)
	}

	return &response
}
