package createpost

import (
	"context"
	"log"
	"pajarit-feed-service/application/ports"
	"pajarit-feed-service/domain"
)

type CreatePost struct {
	postRepository domain.PostRepository
	eventPublisher ports.EventPublisher
}

func NewCreatePost(postRepository domain.PostRepository, eventPublisher ports.EventPublisher) CreatePost {
	return CreatePost{
		postRepository: postRepository,
		eventPublisher: eventPublisher,
	}
}

func (e *CreatePost) Exec(ctx context.Context, cmd CreatePostCmd) (*CreatePostResponse, error) {

	postToSave, err := domain.NewPost(cmd.AuthorId, cmd.Content)
	if err != nil {
		// Se emplea una estrategia de logueo simple a modo ilustrativo
		// pero dependiendo las necesidades debería cambiar
		// (sobre todo si existen integraciones contra DataDog, por ejemplo)
		log.Println(err)
		return nil, err
	}

	savedPost, err := e.postRepository.Save(ctx, postToSave)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	response := NewCreatePostResponse(savedPost)

	go e.eventPublisher.Publish("post.created", response)

	return response, nil
}
