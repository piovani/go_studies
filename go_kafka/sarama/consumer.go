package sarama

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ready chan bool
}

const (
	Broker string = "127.0.0.1:9092"
	Group  string = "0"
	Topic  string = "topic-test"
)

func GetMessagesFromTopic() bool {
	config := sarama.NewConfig()
	config.Version = sarama.DefaultVersion
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.ChannelBufferSize = 2000

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := sarama.NewConsumerGroup(strings.Split(Broker, ","), Group, config)
	if err != nil {
		fmt.Println("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			if err := client.Consume(ctx, strings.Split(Topic, ","), &consumer); err != nil {
				fmt.Println("DEU RUIM, ERRO: %v", err)
			}

			if ctx.Err() != nil {
				return
			}

			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready
	fmt.Println("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		fmt.Println("terminating: context cancelled")

	case <-sigterm:
		fmt.Println("terminating: via signal")
	}

	cancel()
	wg.Wait()

	if err = client.Close(); err != nil {
		fmt.Printf("Error closing client: %v", err)
	}

	return true
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Println(
			"Message claimed: value = %s, timestamp = %v, topic = %s",
			string(message.Value),
			message.Timestamp,
			message.Topic,
		)

		session.MarkMessage(message, "")
	}

	return nil
}
