package ports

type EventPublisher interface {
	Publish(subject string, event any) error
}
