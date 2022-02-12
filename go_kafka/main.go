package main

import (
	"fmt"
	"go_kafka/confluent"

	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
)

// "go_kafka/sarama"

func main() {
	// SARAMA
	// sarama.GetMessagesFromTopic()
	// sarama.PostmessageInTopic()

	// CONFLUENT
	// consumer := confluent.Consumer{}
	// consumer.Consumer()

	var msgChan = make(chan *ckafka.Message)
	configMapConsumer := &ckafka.ConfigMap{
		"bootstrap.servers": "127.0.0.1:9092",
		"group.id":          "appgo",
	}

	topics := []string{"courses"}

	consumer := confluent.NewConsumer(configMapConsumer, topics)

	go consumer.Consumer(msgChan)

	for msg := range msgChan {
		fmt.Println("NOVA MENSASAGE!!!")
		fmt.Println(msg)
	}
}
