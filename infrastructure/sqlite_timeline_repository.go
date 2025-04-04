package infrastructure

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"pajarit-feed-service/domain"
)

type SqliteTimelineRepository struct {
	dbClient *sql.DB
}

func NewSqliteTimelineRepository(dbClient *sql.DB) *SqliteTimelineRepository {
	return &SqliteTimelineRepository{dbClient: dbClient}
}

func (r *SqliteTimelineRepository) Get(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error) {

	// Para el challenge se está usando una base SQLite, tratando de simular el comportamiento
	// de una base key-value, aunque con limitaciones: se guarda un JSON y no una colección.
	// Aun así es posible simular la recuperación de datos de un modo "similar".
	rows, err := r.dbClient.Query("SELECT value AS post FROM timelines, json_each(timelines.posts) WHERE user_id = ? LIMIT ? OFFSET ?", userId, size, offset)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		var postJSON string
		if err := rows.Scan(&postJSON); err != nil {
			return nil, fmt.Errorf("error when reading row: %v", err)
		}

		// Formalmente hablando, los datos recuperados deberían modelarse primero en un DTO y luego
		// adecuarlos a la entidad de negocio.
		// Por cuestiones de practicidad, modelo directamente la entidad.
		var post domain.Post
		if err := json.Unmarshal([]byte(postJSON), &post); err != nil {
			return nil, fmt.Errorf("error when deserializing row: %v", err)
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %v", err)
	}

	timeline := &domain.Timeline{Posts: posts}
	return timeline, nil
}
