package gettimeline

import (
	"context"
	"log"
	"pajarit-feed-service/domain"
)

type GetTimeline struct {
	timelineRepository domain.TimelineRepository
}

func NewGetTimeline(timelineRepository domain.TimelineRepository) GetTimeline {
	return GetTimeline{timelineRepository: timelineRepository}
}

func (e *GetTimeline) Exec(ctx context.Context, cmd GetTimelineCmd) (*TimelineResponse, error) {

	userId := cmd.UserId // TODO - falta aplicar validación

	timeline, err := e.timelineRepository.Get(ctx, userId, cmd.Offset, cmd.Size)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return NewGetTimelineResponse(timeline, cmd), nil
}
