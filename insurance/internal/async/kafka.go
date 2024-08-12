package async

import (
	"context"
	"encoding/json"
	"github.com/IBM/sarama"
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
	defer prod.Close()

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
