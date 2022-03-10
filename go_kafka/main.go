package main

import "go_kafka/confluent"

// samaraLocal "go_kafka/sarama"

func main() {
	// SARAMA
	// samaraLocal := samaraLocal.NewSaramaLocal([]string{"127.0.0.1:9092"}, "topic-test")
	// samaraLocal.GetMessagesFromTopic()
	// samaraLocal.PostmessageInTopic("Eu sou a messagem 3")

	// CONFLUENT
	confluentLocal := confluent.NewConflientLocal()
	confluentLocal.GetMessagesFromTopic()
}
