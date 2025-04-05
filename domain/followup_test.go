package domain

import (
	"testing"
)

func TestNewFollowUp(t *testing.T) {
	tests := []struct {
		name        string
		followerId  string
		followedId  string
		expectError bool
	}{
		{
			name:        "OK: followup válido",
			followerId:  "user-1",
			followedId:  "user-2",
			expectError: false,
		},
		{
			name:        "ERROR: followerId vacío",
			followerId:  "",
			followedId:  "user-2",
			expectError: true,
		},
		{
			name:        "ERROR: followedId vacío",
			followerId:  "user-1",
			followedId:  "",
			expectError: true,
		},
		{
			name:        "ERROR: ambos campos vacíos",
			followerId:  "",
			followedId:  "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := NewFollowUp(tt.followerId, tt.followedId)
			if tt.expectError && err == nil {
				t.Errorf("esperaba error pero fue nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("no se esperaba error pero fue: %v", err)
			}
			if !tt.expectError && (f.FollowerId != tt.followerId || f.FollowedId != tt.followedId) {
				t.Errorf("resultado inesperado: %+v", f)
			}
		})
	}
}
