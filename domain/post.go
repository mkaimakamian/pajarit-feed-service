package domain

import (
	"context"
	"time"
)

type PostRepository interface {
	Save(ctx context.Context, post Post) error
}

// Representa un post muy básico; dependiendo la necesidad del negocio, podrían ser útiles
// fecha de edición, estado de eliminación, cantidad de likes y veces que se compartió.
type Post struct {
	Id           string
	AuthorId     string
	Message      string
	CreationDate time.Time
}
