package config

import (
	"database/sql"
	"fmt"
	"pajarit-feed-service/domain"
	"pajarit-feed-service/infrastructure"

	_ "modernc.org/sqlite"
)

type Dependencies struct {
	PostRepository     domain.PostRepository
	FollowUpRepository domain.FollowUpRepository
	TimelineRepository domain.TimelineRepository
}

func BuildDependencies(cfg *Configuration) (*Dependencies, error) {

	dbClient, err := DBClient(cfg)
	if err != nil {
		return nil, err
	}

	postRepository := infrastructure.NewSqlitePostRepository(dbClient)
	followUpRepository := infrastructure.NewSqliteFollowUpRepository(dbClient)

	deps := &Dependencies{
		FollowUpRepository: followUpRepository,
		PostRepository:     postRepository,
	}

	return deps, nil
}

func DBClient(cfg *Configuration) (*sql.DB, error) {
	client, err := sql.Open("sqlite", cfg.DBPath)
	if err != nil {
		return nil, fmt.Errorf("db can't be opened: %v", err)
	}

	// Valores arbitrarios para el challenge
	client.SetMaxOpenConns(cfg.DBMaxConnection)
	client.SetMaxIdleConns(cfg.DBMaxIdleConnection)

	if err = client.Ping(); err != nil {
		return nil, fmt.Errorf("db is not responding: %v", err)
	}

	return client, nil
}
