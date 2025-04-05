package domain

import (
	"testing"
)

func TestNewPost(t *testing.T) {
	tests := []struct {
		name        string
		authorId    string
		content     string
		expectError bool
	}{
		{
			name:        "OK: post válido",
			authorId:    "user-1",
			content:     "Este es un post válido",
			expectError: false,
		},
		{
			name:        "ERROR: authorId vacío",
			authorId:    "",
			content:     "Contenido válido",
			expectError: true,
		},
		{
			name:        "ERROR: content vacío",
			authorId:    "user-1",
			content:     "",
			expectError: true,
		},
		{
			name:        "ERROR: content demasiado largo",
			authorId:    "user-1",
			content:     makeString(MAX_ALLOWED_LENGTH + 1),
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			post, err := NewPost(tt.authorId, tt.content)

			if tt.expectError && err == nil {
				t.Errorf("esperaba error pero fue nil")
			}
			if !tt.expectError && err != nil {
				t.Errorf("no se esperaba error pero fue: %v", err)
			}
			if !tt.expectError && (post.AuthorId != tt.authorId || post.Content != tt.content) {
				t.Errorf("resultado inesperado: %+v", post)
			}
		})
	}
}

func makeString(n int) string {
	s := make([]rune, n)
	for i := range s {
		s[i] = 'a'
	}
	return string(s)
}
