package kafka

import (
	"github.com/segmentio/kafka-go"
	"context"
	"strconv"
	"log"
	"os"
)

const (
	topic         = "aggregation_actions"
	brokerAddress = "localhost:29092"
)


type Writer struct {
	Writer *kafka.Writer
}

func NewKafkaWriter() *Writer {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	// intialize the writer with the broker addresses, and the topic
	writer := &kafka.Writer{
		Addr: kafka.TCP(brokerAddress),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
		Logger: l,
		RequiredAcks: 1,
	}

	return &Writer{
		Writer: writer,
	}

}

func (w *Writer) WriteMessage (ctx context.Context, data string) error {
	i := 0
	err := w.Writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte(data),
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	
	if err := w.Writer.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}

	return err
}