package async

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/car_insurance/contract/internal/logger"
)

func newKafka[T any](addr, topic string) Consumer[T] {
	consCfg := sarama.NewConfig()

	cons, err := sarama.NewConsumer([]string{addr}, consCfg)
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}
	defer cons.Close()

	partitions, _ := cons.Partitions(topic)

	pcons, err := cons.ConsumePartition(topic, partitions[0], sarama.OffsetNewest)
	if nil != err {
		panic("failed to consume: " + err.Error())
	}

	return &consumer[T]{
		cons:  cons,
		pCons: pcons,
	}
}

type consumer[T any] struct {
	cons  sarama.Consumer
	pCons sarama.PartitionConsumer
}

func (c consumer[T]) Consume(ctx context.Context) <-chan T {
	resChan := make(chan T)
	l := logger.ExtractLogger(ctx)

	defer func() {
		_ = c.cons.Close()
		close(resChan)
	}()

	go func() {
		for {
			select {
			case msg := <-c.pCons.Messages():
				var val T
				if err := json.Unmarshal(msg.Value, &val); err != nil {
					l.Error("failed to unmarshal message", logger.WithArg("error", err.Error()))
					continue
				}

				resChan <- val
			case <-ctx.Done():
				return
			}
		}
	}()

	return resChan
}
