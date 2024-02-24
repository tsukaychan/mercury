package events

import (
	"encoding/json"

	"github.com/IBM/sarama"
)

const topicReadEvent = "article_read_event"

type ReadEvent struct {
	Aid int64
	Uid int64
}

type SaramaSyncProducer struct {
	producer sarama.SyncProducer
}

func NewSaramaSyncProducer(producer sarama.SyncProducer) Producer {
	return &SaramaSyncProducer{
		producer: producer,
	}
}

func (pdr *SaramaSyncProducer) ProduceReadEvent(evt ReadEvent) error {
	val, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	_, _, err = pdr.producer.
		SendMessage(&sarama.ProducerMessage{
			Topic: topicReadEvent,
			Value: sarama.ByteEncoder(val),
		})
	return err
}
