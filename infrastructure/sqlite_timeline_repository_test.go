package infrastructure

import (
	"context"
	"database/sql"
	"testing"
)

func setupTimelineTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory DB: %v", err)
	}

	schema := `
		CREATE TABLE timelines (
			user_id TEXT PRIMARY KEY,
			posts TEXT
		);`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return db
}

func TestSqliteTimelineRepository_Get(t *testing.T) {
	db := setupTimelineTestDB(t)
	defer db.Close()

	repo := NewSqliteTimelineRepository(db)

	// Timeline fake con 2 posts
	jsonPosts := `[{
		"id": "post-1",
		"author_id": "user-123",
		"content": "Primer post",
		"created_at": "2024-01-01T10:00:00Z"
	},
	{
		"id": "post-2",
		"author_id": "user-456",
		"content": "Segundo post",
		"created_at": "2024-01-02T10:00:00Z"
	}]`

	_, err := db.Exec(`INSERT INTO timelines (user_id, posts) VALUES (?, ?)`, "user-999", jsonPosts)
	if err != nil {
		t.Fatalf("failed to insert timeline: %v", err)
	}

	timeline, err := repo.Get(context.Background(), "user-999", 0, 10)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(timeline.Posts) != 2 {
		t.Errorf("expected 2 posts, got %d", len(timeline.Posts))
	}

	if timeline.Posts[0].Id != "post-1" || timeline.Posts[1].Id != "post-2" {
		t.Error("posts not returned in expected order or with correct IDs")
	}
}
