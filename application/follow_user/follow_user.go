package followuser

import (
	"context"
	"log"
	"pajarit-feed-service/domain"
)

type FollowUser struct {
	followUpRepository domain.FollowUpRepository
}

func NewFollowUser(followUpRepository domain.FollowUpRepository) FollowUser {
	return FollowUser{followUpRepository: followUpRepository}
}

func (e *FollowUser) Exec(ctx context.Context, cmd FollowUsertCmd) error {

	followUp, err := domain.NewFollowUp(cmd.FollowerId, cmd.FollowedId)
	if err != nil {
		log.Println(err)
		return err
	}

	err = e.followUpRepository.Save(ctx, followUp)
	if err != nil {
		log.Println(err)
		return err
	}

	// En este caso no aporta valor devolver la representación de la relación
	return nil
}
