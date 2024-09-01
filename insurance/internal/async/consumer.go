package async

import "context"

type Consumer[T any] interface {
	Consume(ctx context.Context) chan T
	Close() error
}

func NewConsumer[T any](consType uint, addr, topic string) Consumer[T] {
	switch consType {
	case Kafka:
		return newKafkaConsumer[T](addr, topic)
	default:
		panic("invalid consumer type")
	}
}
