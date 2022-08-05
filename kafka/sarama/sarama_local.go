package sarama

type SaramaLocal struct {
	Brokers []string
	Topic   string
}

func NewSaramaLocal(brokers []string, topic string) SaramaLocal {
	return SaramaLocal{
		Brokers: brokers,
		Topic:   topic,
	}
}
