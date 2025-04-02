package domain

import (
	"context"
	"fmt"
)

type FollowUpRepository interface {
	Save(ctx context.Context, followUp *FollowUp) error
}

// Representa una relación muy básica de seguidor - seguido.
// Idealmente debería contar con un atributo que de idea de la fecha de alta;
// otros datos de auditoría también podrían ser útiles a menos que se audite por otro medio.
type FollowUp struct {
	FollowerId string
	FollowedId string
}

func NewFollowUp(followerId, followedId string) (*FollowUp, error) {

	// Podemos desacoplar la validación, pero me resultó más
	// práctico tratar la entidad como un value object

	if len(followerId) == ZERO_LENGTH {
		return nil, fmt.Errorf("follower id can't be %d length", ZERO_LENGTH)
	}

	if len(followedId) == ZERO_LENGTH {
		return nil, fmt.Errorf("followed id can't be %d length", ZERO_LENGTH)
	}

	return &FollowUp{FollowerId: followerId, FollowedId: followedId}, nil
}
