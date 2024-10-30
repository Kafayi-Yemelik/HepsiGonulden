package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/viper"
)

type Producer struct {
	producer sarama.SyncProducer
}

func NewProducer() *Producer {
	brokers := viper.GetStringSlice("kafka.brokers")

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}

	return &Producer{
		producer: producer,
	}
}

func (p *Producer) Publish(topic string, value any) error {
	message := &sarama.ProducerMessage{Topic: topic}

	jsonValue, err := json.Marshal(&value)
	if err != nil {
		return err
	}

	message.Value = sarama.ByteEncoder(jsonValue)

	partition, offset, err := p.producer.SendMessage(message)
	if err != nil {
		return err
	}

	fmt.Printf("topic=%s\tpartition=%d\toffset=%d\n", topic, partition, offset)
	return nil
}
