package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"pajarit-feed-service/domain"
)

type SqliteFollowUpRepository struct {
	dbClient *sql.DB
}

func NewSqliteFollowUpRepository(dbClient *sql.DB) *SqliteFollowUpRepository {
	return &SqliteFollowUpRepository{dbClient: dbClient}
}

func (r *SqliteFollowUpRepository) Save(ctx context.Context, followUp *domain.FollowUp) error {

	_, err := r.dbClient.Exec("INSERT INTO followup (follower_id, followed_id) VALUES (?, ?)", followUp.FollowerId, followUp.FollowedId)
	if err != nil {
		return fmt.Errorf("can't insert followup relation %v", err)
	}

	return nil
}
