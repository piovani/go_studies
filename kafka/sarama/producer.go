package sarama

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func (s *SaramaLocal) PostmessageInTopic(message string) {
	producer, err := s.newProducer()
	if err != nil {
		fmt.Println("Could not create producer: ", err)
	}

	msg := prepareMessage(s.Topic, message)

	partition, offset, err := producer.SendMessage(msg)

	fmt.Println(partition, offset, err)

	if err != nil {
		fmt.Sprintln("%s error occured.", err.Error())
	} else {
		fmt.Sprintln("Message was saved to partion: %d.\nMessage offset is: %d.\n", partition, offset)
	}
}

func (s *SaramaLocal) newProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(s.Brokers, config)

	return producer, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}

	return msg
}
