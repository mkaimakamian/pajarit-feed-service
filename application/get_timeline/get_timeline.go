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

	// savePost, err := domain.NewPost(cmd.AuthorId, cmd.Message)
	// if err != nil {
	// 	// Se emplea una estrategia de logueo simple a modo ilustrativo
	// 	// pero dependiendo las necesidades debería cambiar
	// 	// (sobre todo si existen integraciones contra DataDog, por ejemplo)
	// 	log.Println(err)
	// 	return nil, err
	// }

	userId := cmd.UserId // TODO - falta aplicar validación

	timeline, err := e.timelineRepository.Get(ctx, userId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return NewGetTimelineResponse(timeline, userId), nil
}
