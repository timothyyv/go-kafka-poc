package operations


import (
	"fmt"
	"context"
	// "strconv"
	"math/rand"
	// "grule-demo/kafka"
	"grule-demo/grpc"
	// kafkago "github.com/segmentio/kafka-go"
	"github.com/hyperjumptech/grule-rule-engine/logger"
)

type CustomFunction struct {}

func (cf *CustomFunction) Aggregation(input string) float64 {
	ctx := context.Background()
	connector := grpc.NewGRPCConnector()
	res, err := connector.SendMessage(ctx, input)
	if err != nil {
		logger.Log.Error("RESPONSE ERROR", err)
	}
	// reader := kafka.NewKafkaReader([]string{"sum_response"})
	// writer := kafka.NewKafkaWriter()

	// ctx := context.Background()
	// messages := make(chan kafkago.Message, 1000)

	// go reader.FetchMessages(ctx, messages)

	// go writer.WriteMessage(ctx, input)

	// go reader.CommitMessages(ctx, messages)

	// x := <- messages

	// fmt.Printf("%v 9283003", string(x.Value))
	// y,err := strconv.ParseInt(string(x.Value), 10, 64)
	// if err != nil {
    //     fmt.Printf("%d of type %T", y, y)
    // }
	fmt.Printf("Response from GRPC: %v", res.Value)
	if res.Value <= 0 {
		return 9999999999999.999
	}
	return float64(res.Value)
}

func makeTopic() string {
	return fmt.Sprintf("kafka-go-%016x", rand.Int63())
}

func makeGroupID() string {
	return fmt.Sprintf("kafka-go-group-%016x", rand.Int63())
}
