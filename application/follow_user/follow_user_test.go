package followuser

import (
	"context"
	"errors"
	"testing"

	"pajarit-feed-service/domain"
)

// ------------------------
type MockFollowUpRepository struct {
	SaveFunc func(ctx context.Context, followUp *domain.FollowUp) error
}

func (m *MockFollowUpRepository) Save(ctx context.Context, followUp *domain.FollowUp) error {
	return m.SaveFunc(ctx, followUp)
}

// ------------------------

func TestFollowUser_Exec(t *testing.T) {
	tests := []struct {
		name         string
		cmd          FollowUsertCmd
		mockSaveFunc func(ctx context.Context, followUp *domain.FollowUp) error
		expectError  bool
	}{
		{
			name: "Relación válida",
			cmd: FollowUsertCmd{
				FollowerId: "user1",
				FollowedId: "user2",
			},
			mockSaveFunc: func(ctx context.Context, followUp *domain.FollowUp) error {
				return nil
			},
			expectError: false,
		},
		{
			name: "FollowerId vacío",
			cmd: FollowUsertCmd{
				FollowerId: "",
				FollowedId: "user2",
			},
			mockSaveFunc: nil,
			expectError:  true,
		},
		{
			name: "FollowedId vacío",
			cmd: FollowUsertCmd{
				FollowerId: "user1",
				FollowedId: "",
			},
			mockSaveFunc: nil,
			expectError:  true,
		},
		{
			name: "Error en repo",
			cmd: FollowUsertCmd{
				FollowerId: "user1",
				FollowedId: "user2",
			},
			mockSaveFunc: func(ctx context.Context, followUp *domain.FollowUp) error {
				return errors.New("fallo en persistencia")
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockFollowUpRepository{
				SaveFunc: func(ctx context.Context, followUp *domain.FollowUp) error {
					if tt.mockSaveFunc != nil {
						return tt.mockSaveFunc(ctx, followUp)
					}
					t.Fatal("Save() no debería haberse llamado")
					return nil
				},
			}

			usecase := NewFollowUser(mockRepo)

			err := usecase.Exec(context.Background(), tt.cmd)

			if tt.expectError && err == nil {
				t.Errorf("esperaba error pero fue nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("no esperaba error pero fue: %v", err)
			}
		})
	}
}
