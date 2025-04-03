package infrastructure

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)

type NatsEventPublisher struct {
	conn *nats.Conn
}

func NewNatsEventPublisher(conn *nats.Conn) *NatsEventPublisher {
	return &NatsEventPublisher{conn: conn}
}

func (p *NatsEventPublisher) Publish(subject string, event any) error {
	data, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return p.conn.Publish(subject, data)
}
