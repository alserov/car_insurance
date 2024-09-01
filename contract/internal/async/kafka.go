package async

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/alserov/car_insurance/contract/internal/logger"
	"github.com/alserov/car_insurance/contract/internal/utils"
)

func newKafka[T any](addr, topic string) Consumer[T] {
	consCfg := sarama.NewConfig()

	cons, err := sarama.NewConsumer([]string{addr}, consCfg)
	if err != nil {
		panic("failed to init consumer: " + err.Error())
	}

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

func (c consumer[T]) Close() error {
	_ = c.cons.Close()
	_ = c.pCons.Close()
	return nil
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
