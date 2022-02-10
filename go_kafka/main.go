package main

import (
	"go_kafka/confluent"
	// "go_kafka/sarama"
)

func main() {
	// SARAMA
	// sarama.GetMessagesFromTopic()
	// sarama.PostmessageInTopic()

	// CONFLUENT
	consumer := confluent.Consumer{}
	consumer.Consumer()
}
