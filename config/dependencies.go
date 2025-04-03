package config

import (
	"database/sql"
	"fmt"
	"pajarit-feed-service/domain"

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

	// Crear instancia de los repos
	fmt.Print(dbClient)
	deps := &Dependencies{}
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
