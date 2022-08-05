package sarama

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func (s *SaramaLocal) GetMessagesFromTopic() {
	consumer, err := sarama.NewConsumer(s.Brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	subscribe(s.Topic, consumer)
}

func subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Erro ao listar as particoes ", err)
	}
	initialOffset := sarama.OffsetOldest

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	fmt.Println(string(message.Value))
}
