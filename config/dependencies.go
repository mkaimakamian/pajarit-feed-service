package config

import "pajarit-feed-service/domain"

type Dependencies struct {
	PostRepository     domain.PostRepository
	FollowUpRepository domain.FollowUpRepository
	TimelineRepository domain.TimelineRepository
}

func BuildDependencies(config *Configuration) (*Dependencies, error) {

	deps := &Dependencies{}
	return deps, nil
}
