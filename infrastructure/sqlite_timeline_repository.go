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

// func (r *SqliteTimelineRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {

// 	toInsert := post
// 	toInsert.Id = uuid.New().String()
// 	toInsert.CreatedAt = time.Now().UTC()

// 	_, err := r.dbClient.Exec(
// 		"INSERT INTO posts (id, author_id, content, created_at) VALUES (?, ?, ?, ?)",
// 		toInsert.Id, post.AuthorId, post.Content, post.CreatedAt,
// 	)

// 	// TODO - tipar el error

// 	if err != nil {
// 		return nil, fmt.Errorf("can't insert post %v", err)
// 	}

// 	return toInsert, nil
// }

func (r *SqliteTimelineRepository) Get(ctx context.Context, userId string) (*domain.Timeline, error) {

	rows, err := r.dbClient.Query("SELECT value AS post FROM timelines, json_each(timelines.posts) WHERE user_id = ? LIMIT 10 OFFSET ?", userId, 0)

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
