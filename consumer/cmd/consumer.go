package cmd

import (
	"context"
	"errors"
	"github.com/IBM/sarama"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Consumer struct {
	ready chan bool
}

type GlobalErrorHandlerResp struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ConsumerOrderCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "consumer",
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

			config := sarama.NewConfig()
			config.Version = version
			config.Consumer.Offsets.Initial = sarama.OffsetOldest

			consumer := Consumer{
				ready: make(chan bool),
			}

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

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				log.Printf("message channel was closed")
				return nil
			}
			log.Printf("Message claimed: value = %s, timestamp = %v, topic = %s", string(message.Value), message.Timestamp, message.Topic)
			session.MarkMessage(message, "")
		case <-session.Context().Done():
			return nil
		}
	}
}
