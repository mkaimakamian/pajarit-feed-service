package createpost

import (
	"context"
	"log"
	"pajarit-feed-service/domain"
)

type CreatePost struct {
	postRepository domain.PostRepository
}

func NewCreatePost() CreatePost {
	return CreatePost{}
}

func (e *CreatePost) Exec(ctx context.Context, cmd CreatePostCmd) (*CreatePostResponse, error) {

	savePost, err := domain.NewPost(cmd.AuthorId, cmd.Message)
	if err != nil {
		// Se emplea una estrategia de logueo simple a modo ilustrativo
		// pero dependiendo las necesidades debería cambiar
		// (sobre todo si existen integraciones contra DataDog, por ejemplo)
		log.Println(err)
		return nil, err
	}

	saved, err := e.postRepository.Save(ctx, savePost)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 4. Disparar el evento para su posterior réplica
	// TODO - Implementar los eventos

	return NewCreatePostResponse(saved), nil
}
