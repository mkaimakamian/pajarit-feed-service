package gettimeline

import (
	"context"
	"errors"
	"testing"
	"time"

	"pajarit-feed-service/domain"
)

// ------------------------
type MockTimelineRepository struct {
	GetFunc func(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error)
}

func (m *MockTimelineRepository) Get(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error) {
	return m.GetFunc(ctx, userId, offset, size)
}

// ------------------------

func TestGetTimeline_Exec(t *testing.T) {
	tests := []struct {
		name        string
		cmd         GetTimelineCmd
		mockGetFunc func(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error)
		expectErr   bool
		expectFeed  int
	}{
		{
			name: "Timeline recuperado con éxito",
			cmd: GetTimelineCmd{
				UserId: "user123",
				Offset: 0,
				Size:   2,
			},
			mockGetFunc: func(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error) {
				return &domain.Timeline{
					Posts: []domain.Post{
						{Id: "p1", AuthorId: "a1", Content: "Aguante", CreatedAt: time.Now()},
						{Id: "p2", AuthorId: "a2", Content: "Pájarit!", CreatedAt: time.Now()},
					},
				}, nil
			},
			expectErr:  false,
			expectFeed: 2,
		},
		{
			name: "Error en repositorio",
			cmd: GetTimelineCmd{
				UserId: "user123",
				Offset: 0,
				Size:   10,
			},
			mockGetFunc: func(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error) {
				return nil, errors.New("fallo en el repo")
			},
			expectErr:  true,
			expectFeed: 0,
		},
		{
			name: "Timeline vacío",
			cmd: GetTimelineCmd{
				UserId: "user123",
				Offset: 0,
				Size:   10,
			},
			mockGetFunc: func(ctx context.Context, userId string, offset, size int) (*domain.Timeline, error) {
				return &domain.Timeline{
					Posts: []domain.Post{},
				}, nil
			},
			expectErr:  false,
			expectFeed: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTimelineRepository{
				GetFunc: tt.mockGetFunc,
				// Como no implementé la validación del usuario (to do!)
				// en todos los casos voy a estar llamando a la implementación
				// del mock y no hace falta que controle si se llama inadvertidamente
			}

			usecase := NewGetTimeline(mockRepo)

			resp, err := usecase.Exec(context.Background(), tt.cmd)

			if tt.expectErr && err == nil {
				t.Errorf("esperaba error pero fue nil")
			}
			if !tt.expectErr && err != nil {
				t.Errorf("no esperaba error pero fue: %v", err)
			}
			if err == nil && resp != nil && len(resp.Feed) != tt.expectFeed {
				t.Errorf("esperaba %d posts en el feed pero obtuve %d", tt.expectFeed, len(resp.Feed))
			}
		})
	}
}
