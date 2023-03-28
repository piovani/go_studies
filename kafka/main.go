package main

import (
	confluentLocal "go_kafka/confluent"
	samaraLocal "go_kafka/sarama"
)

var (
	SaramaLocal    = samaraLocal.NewSaramaLocal([]string{"127.0.0.1:9092"}, "topic-test")
	ConfluentLocal = confluentLocal.NewConflientLocal()
)

func main() {
	// SARAMA
	// SaramaLocal.GetMessagesFromTopic()
	SaramaLocal.PostmessageInTopic("Eu sou a messagem 3")

	// CONFLUENT
	// ConfluentLocal.GetMessagesFromTopic()
}
