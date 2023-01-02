package kafka

import (
	"github.com/segmentio/kafka-go"
	"github.com/pkg/errors"
	"context"
	"log"
	"os"
)


type Reader struct {
	Reader *kafka.Reader
}

func NewKafkaReader(topics []string) *Reader {
	l := log.New(os.Stdout, "kafka reader: ", 0)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{brokerAddress},
		// Topic:   topic,
		GroupID: "group-key",
		GroupTopics: topics,
		// assign the logger to the reader
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		StartOffset: kafka.LastOffset,
		Logger: l,
	})

	return &Reader{
		Reader: reader,
	}
}

func (k *Reader) FetchMessages(ctx context.Context, messages chan<- kafka.Message) error {
	for {
		message, err := k.Reader.FetchMessage(ctx)
		if err != nil {
			return err
		}

		// check if context is still active
		select {
		case <- ctx.Done():
			return ctx.Err()
		case messages <- message:
			log.Printf("message fetched and sent to a channel: %v \n", string(message.Value))
		}
	}
}

func (k *Reader) CommitMessages(ctx context.Context, messages <-chan kafka.Message) error {
	for {
		select {
		case <- ctx.Done():
		case msg := <- messages:
			err := k.Reader.CommitMessages(ctx, msg)
			if err != nil {
				return errors.Wrap(err, "Reader.CommitMessages")
			}
			log.Printf("committed a msg: %v \n", string(msg.Value))
		}
	}
}