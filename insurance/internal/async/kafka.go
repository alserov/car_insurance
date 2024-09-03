package async

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/car_insurance/insurance/internal/logger"
	"github.com/alserov/car_insurance/insurance/internal/utils"
)

func newKafkaProducer(addr, topic string) Producer {
	prodCfg := sarama.NewConfig()
	prodCfg.Producer.Partitioner = sarama.NewRandomPartitioner
	prodCfg.Producer.RequiredAcks = sarama.WaitForAll
	prodCfg.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer([]string{addr}, prodCfg)
	if err != nil {
		panic("failed to init producer: " + err.Error())
	}

	return &producer{
		topic: topic,
		p:     prod,
	}
}

type producer struct {
	topic string

	p sarama.SyncProducer
}

func (p producer) Close() error {
	return p.p.Close()
}

func (p producer) Produce(ctx context.Context, val any) error {
	b, err := json.Marshal(val)
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	_, _, err = p.p.SendMessage(&sarama.ProducerMessage{
		Topic: p.topic,
		Value: sarama.StringEncoder(b),
	})
	if err != nil {
		return utils.NewError(err.Error(), utils.Internal)
	}

	return nil
}

func newKafkaConsumer[T any](addr, topic string) Consumer[T] {
	cfg := sarama.NewConfig()

	cons, err := sarama.NewConsumer([]string{addr}, cfg)
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

	return &consumer[T]{
		c:     cons,
		topic: topic,
	}
}

type consumer[T any] struct {
	c sarama.Consumer

	topic string
}

func (c consumer[T]) Close() error {
	return c.c.Close()
}

func (c consumer[T]) Consume(ctx context.Context) chan T {
	partitions, _ := c.c.Partitions(c.topic)

	// consuming partition
	pcons, err := c.c.ConsumePartition(c.topic, partitions[0], sarama.OffsetNewest)
	if nil != err {
		panic("failed to consume: " + err.Error())
	}

	chRes := make(chan T)

	go func() {
		defer close(chRes)

		for {
			select {
			case msg := <-pcons.Messages():
				var val T
				if err = json.Unmarshal(msg.Value, &val); err != nil {
					logger.ExtractLogger(ctx).Error("failed to unmarshal message", logger.WithArg("error", err.Error()))
					continue
				}

				chRes <- val
			case <-ctx.Done():
				return
			}
		}
	}()

	return chRes
}
