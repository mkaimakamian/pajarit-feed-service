package infrastructure

import (
	"context"
	"database/sql"
	"pajarit-feed-service/domain"
	"testing"

	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory DB: %v", err)
	}

	schema := `
		CREATE TABLE posts (
			id TEXT PRIMARY KEY,
			author_id TEXT,
			content TEXT,
			created_at TEXT
		);`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return db
}

func TestSqlitePostRepository_Save(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := NewSqlitePostRepository(db)

	post := &domain.Post{
		AuthorId: "user-123",
		Content:  "Qué grande Pájarit! mejor que Twitter!",
	}

	savedPost, err := repo.Save(context.Background(), post)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if savedPost.Id == "" {
		t.Error("expected post to have an ID")
	}

	if savedPost.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}

	// Controlo que se haya insertado algo... simepre es bueno :)
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM posts WHERE id = ?", savedPost.Id).Scan(&count)
	if err != nil || count != 1 {
		t.Error("expected post to be saved in DB")
	}
}
