package gettimeline

import (
	"pajarit-feed-service/domain"
	"time"
)

type GetTimelineCmd struct {
	UserId string
	Offset int
	Size   int
}

type TimelineResponse struct {
	UserId string
	Feed   []PostResponse

	Offset int
	Size   int
}

// Si bien podría haber referenciado a a createpost.CreatePostResponse, previa
// reubicación en un paquete común, prioricé el aislamiento por resultar más
// práctico para el challenge.
type PostResponse struct {
	Id        string
	AuthorId  string
	Content   string
	CreatedAt time.Time
}

func NewGetTimelineResponse(timeline *domain.Timeline, cmd GetTimelineCmd) *TimelineResponse {

	response := TimelineResponse{UserId: cmd.UserId}
	response.Offset = cmd.Offset
	response.Size = len(timeline.Posts)

	for _, post := range timeline.Posts {
		post := PostResponse(post)
		response.Feed = append(response.Feed, post)
	}

	return &response
}
