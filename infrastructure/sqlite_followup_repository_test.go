package infrastructure

import (
	"context"
	"database/sql"
	"pajarit-feed-service/domain"
	"testing"

	_ "modernc.org/sqlite"
)

func setupFollowUpTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("failed to open in-memory DB: %v", err)
	}

	schema := `
		CREATE TABLE followup (
			follower_id TEXT,
			followed_id TEXT,
			PRIMARY KEY (follower_id, followed_id)
		);`

	_, err = db.Exec(schema)
	if err != nil {
		t.Fatalf("failed to create schema: %v", err)
	}

	return db
}

func TestSqliteFollowUpRepository_Save(t *testing.T) {
	db := setupFollowUpTestDB(t)
	defer db.Close()

	repo := NewSqliteFollowUpRepository(db)

	followUp := &domain.FollowUp{
		FollowerId: "user-a",
		FollowedId: "user-b",
	}

	err := repo.Save(context.Background(), followUp)
	if err != nil {
		t.Fatalf("unexpected error saving followup: %v", err)
	}

	// Verificamos que se haya guardado
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM followup WHERE follower_id = ? AND followed_id = ?", followUp.FollowerId, followUp.FollowedId).Scan(&count)
	if err != nil {
		t.Fatalf("failed to query followup: %v", err)
	}

	if count != 1 {
		t.Errorf("expected 1 followup record, got %d", count)
	}
}
