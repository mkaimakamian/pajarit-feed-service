package config

import (
	"database/sql"
	"fmt"
	"log"
	"pajarit-feed-service/application/ports"
	"pajarit-feed-service/domain"
	"pajarit-feed-service/infrastructure"

	"github.com/nats-io/nats.go"
	_ "modernc.org/sqlite"
)

type Dependencies struct {
	PostRepository     domain.PostRepository
	FollowUpRepository domain.FollowUpRepository
	TimelineRepository domain.TimelineRepository
	EventPublisher     ports.EventPublisher
}

func BuildDependencies(cfg *Configuration) (*Dependencies, error) {

	dbClient, err := dbClient(cfg)
	if err != nil {
		return nil, err
	}

	postRepository := infrastructure.NewSqlitePostRepository(dbClient)
	followUpRepository := infrastructure.NewSqliteFollowUpRepository(dbClient)
	timelineRepository := infrastructure.NewSqliteTimelineRepository(dbClient)

	epConnection := eventPublisherConnection(cfg)
	eventPublisher := infrastructure.NewNatsEventPublisher(epConnection)

	deps := &Dependencies{
		FollowUpRepository: followUpRepository,
		PostRepository:     postRepository,
		TimelineRepository: timelineRepository,
		EventPublisher:     eventPublisher,
	}

	return deps, nil
}

func dbClient(cfg *Configuration) (*sql.DB, error) {
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

func eventPublisherConnection(cfg *Configuration) *nats.Conn {
	natsUrl := fmt.Sprintf("%s:%d", cfg.EventServer, cfg.EventServerPort)
	connection, err := nats.Connect(natsUrl)
	if err != nil {
		log.Fatalf("can't connect to NATS: %v", err)
	}

	return connection
}
