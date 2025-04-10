package consumer

import (
	client2 "HepsiGonulden/client"
	"HepsiGonulden/internal/types"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type OrderCreateConsumer struct {
	ready       chan bool
	orderClient *client2.HttpOrderClient
}

func NewOrderCreateConsumerCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "order-create-consumer",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			brokers := viper.GetStringSlice("kafka.brokers")
			topic := viper.GetString("kafka.topic")
			consumerGroupName := viper.GetString("kafka.consumer_group_name")
			keepRunning := true
			log.Println("Starting a new Sarama consumer")
			sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

			version, err := sarama.ParseKafkaVersion(sarama.V3_0_0_0.String())
			if err != nil {
				log.Panicf("Error parsing Kafka version: %v", err)
			}

			consumer := OrderCreateConsumer{
				ready:       make(chan bool),
				orderClient: client2.NewHttpOrderClient("http://localhost:3001"),
			}

			config := sarama.NewConfig()
			config.Version = version
			config.Consumer.Offsets.Initial = sarama.OffsetOldest

			ctx, cancel := context.WithCancel(context.Background())
			client, err := sarama.NewConsumerGroup(brokers, consumerGroupName, config)
			if err != nil {
				log.Panicf("Error creating consumer group client: %v", err)
			}

			wg := &sync.WaitGroup{}
			wg.Add(1)
			go func() {
				defer wg.Done()
				for {
					if err := client.Consume(ctx, []string{topic}, &consumer); err != nil {
						if errors.Is(err, sarama.ErrClosedConsumerGroup) {
							return
						}
						log.Panicf("Error from consumer: %v", err)
					}
					// check if context was cancelled, signaling that the consumer should stop
					if ctx.Err() != nil {
						return
					}
					consumer.ready = make(chan bool)
				}
			}()
			<-consumer.ready // Await till the consumer has been set up
			log.Println("Sarama consumer up and running!...")

			sigterm := make(chan os.Signal, 1)
			signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

			for keepRunning {
				select {
				case <-ctx.Done():
					log.Println("terminating: context cancelled")
					keepRunning = false
				case <-sigterm:
					log.Println("terminating: via signal")
					keepRunning = false
				}
			}
			cancel()
			wg.Wait()
			if err = client.Close(); err != nil {
				log.Panicf("Error closing client: %v", err)
			}
			return nil
		},
	}

	return rootCmd
}

func (consumer *OrderCreateConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *OrderCreateConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *OrderCreateConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)

			var currentOrder types.Order
			err := json.Unmarshal(message.Value, &currentOrder)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			err = consumer.orderClient.UpdateOrder(context.Background(), currentOrder.Id, types.OrderUpdateModel{
				OrderName:     currentOrder.OrderName,
				OrderTotal:    currentOrder.OrderTotal,
				PaymentMethod: currentOrder.PaymentMethod,
				OrderStatus:   "Ready",
			})
			if err != nil {
				fmt.Printf("order update operation failed, err: %s", err.Error())
				continue
			}

			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
