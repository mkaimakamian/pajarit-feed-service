package infrastructure

import (
	"encoding/json"
	"math"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

const MAX_RETRIES = 6
const BASE_DELAY = time.Millisecond * 100

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

	for i := range MAX_RETRIES {
		err = p.conn.Publish(subject, data)
		if err == nil {
			return nil
		}

		backoff := BASE_DELAY * time.Duration(math.Pow(2, float64(i)))
		jitter := time.Duration(rand.Int63n(int64(backoff / 2)))
		sleepDuration := backoff + jitter
		time.Sleep(sleepDuration)
	}

	return p.conn.Publish(subject, data)
}
