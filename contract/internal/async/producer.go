package async

import "context"

type Producer interface {
	Produce(ctx context.Context, val any) error
	Close() error
}

func NewProducer(prodType int, addr, topic string) Producer {
	switch prodType {
	case Kafka:
		return newKafkaProducer(addr, topic)
	default:
		panic("invalid producer type")
	}
}
