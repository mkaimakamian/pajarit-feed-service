package domain

import "context"

type TimelineRepository interface {
	Get(ctx context.Context, userId string, offset, size int) (*Timeline, error)
}

type Timeline struct {
	Posts []Post
}
