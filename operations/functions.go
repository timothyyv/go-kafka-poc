package operations


import (
	"fmt"
	"context"
	"strconv"
	"grule-demo/kafka"
	kafkago "github.com/segmentio/kafka-go"
)

type CustomFunction struct {}

func (cf *CustomFunction) Aggregation(input string) int64 {
	reader := kafka.NewKafkaReader("sum_response")
	writer := kafka.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafkago.Message, 1000)

	go reader.FetchMessages(ctx, messages)

	go writer.WriteMessage(ctx, input)

	go reader.CommitMessages(ctx, messages)

	x := <- messages

	fmt.Printf("%v 9283003", string(x.Value))
	y,err := strconv.ParseInt(string(x.Value), 10, 64)
	if err != nil {
        fmt.Printf("%d of type %T", y, y)
    }
	return y
}
