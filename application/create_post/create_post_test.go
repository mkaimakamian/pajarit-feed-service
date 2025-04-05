package createpost

import (
	"context"
	"errors"
	"testing"

	"pajarit-feed-service/domain"
)

// ------------------------
type MockPostRepository struct {
	SaveFunc func(ctx context.Context, post *domain.Post) (*domain.Post, error)
}

func (m *MockPostRepository) Save(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	return m.SaveFunc(ctx, post)
}

// ------------------------

// ------------------------
type MockEventPublisher struct {
	PublishFunc func(subject string, event any) error
}

func (m *MockEventPublisher) Publish(subject string, event any) error {
	return m.PublishFunc(subject, event)
}

// ------------------------

func TestCreatePost_Exec(t *testing.T) {
	tests := []struct {
		name         string
		cmd          CreatePostCmd
		mockSaveFunc func(ctx context.Context, post *domain.Post) (*domain.Post, error)
		expectError  bool
	}{
		{
			name: "post válido",
			cmd:  CreatePostCmd{AuthorId: "user1", Content: "Hola mundo! qué lindo es pájarit!"},
			mockSaveFunc: func(ctx context.Context, post *domain.Post) (*domain.Post, error) {
				post.Id = "123ABC"
				return post, nil
			},
			expectError: false,
		},
		{
			name:         "contenido vacío",
			cmd:          CreatePostCmd{AuthorId: "user1", Content: ""},
			mockSaveFunc: nil,
			expectError:  true,
		},
		{
			name: "error en repo",
			cmd:  CreatePostCmd{AuthorId: "user1", Content: "Test"},
			mockSaveFunc: func(ctx context.Context, post *domain.Post) (*domain.Post, error) {
				return nil, errors.New("falló la db")
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockPostRepository{
				SaveFunc: func(ctx context.Context, post *domain.Post) (*domain.Post, error) {
					if tt.mockSaveFunc != nil {
						return tt.mockSaveFunc(ctx, post)
					}
					t.Fatal("Save no debería haberse llamado")
					return nil, nil
				},
			}

			// El publisher se ejecuta en una go routine
			// y sólo loguea en caso de  fallo; no hay mucho
			// para mockear
			mockPublisher := &MockEventPublisher{
				PublishFunc: func(subject string, event any) error {
					return nil
				},
			}

			usecase := CreatePost{
				postRepository: mockRepo,
				eventPublisher: mockPublisher,
			}

			_, err := usecase.Exec(context.Background(), tt.cmd)

			if tt.expectError && err == nil {
				t.Errorf("esperaba error pero no hubo")
			}

			if !tt.expectError && err != nil {
				t.Errorf("no esperaba error pero hubo: %v", err)
			}
		})
	}
}
