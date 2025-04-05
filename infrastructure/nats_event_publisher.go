package infrastructure

import (
	"encoding/json"
	"log"
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

	// En un escenario real, habría que definir adecuadamente
	// los tiempos del exponential backoff, como así también
	// plantear una estrategia sólida en caso extremo de que
	// no se pueda disparar el evento.
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

	if err != nil {
		log.Println(err)
	}

	return err
}
