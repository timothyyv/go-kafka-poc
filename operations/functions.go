package operations


import (
	"context"
	"fmt"

	kafkago "github.com/segmentio/kafka-go"
	// "golang.org/x/sync/errgroup"

	// "fmt"
	// "log"

	// "fmt"
	// "strings"
	"grule-demo/kafka"
	"strconv"
	// "sync"
	// "encoding/json"
	// "log"
)

type CustomFunction struct {}

func (cf *CustomFunction) Aggregation(input string) int64 {
	// var wg sync.WaitGroup

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	kafka.SendToKafka(ctx, input)
	// }()
	// wg.Wait()
	// res, _ := kafka.ConsumeMessage(ctx)
	// fmt.Printf("%v", res)
	// if err != nil {
	// 	log.Fatal("could not read message " + err.Error())
	// }
	// result, err := strconv.ParseInt(res, 10, 64)
	// if err != nil {
	// 	panic(err)
	// }
	reader := kafka.NewKafkaReader("sum_aggregation_response")
	writer := kafka.NewKafkaWriter()

	ctx := context.Background()
	messages := make(chan kafkago.Message, 1000)

	// g, ctx := errgroup.WithContext(ctx)

	// g.Go(func() error {
		go reader.FetchMessages(ctx, messages)
	// })

	// g.Go(func() error {
		go writer.WriteMessage(ctx, input)
	// })


	// g.Go(func() error {
		go reader.CommitMessages(ctx, messages)
	// })

	// err := g.Wait()
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	x := <- messages

	fmt.Printf("%v 9283003", string(x.Value))
	y,err := strconv.ParseInt(string(x.Value), 10, 64)
	if err != nil {
        fmt.Printf("%d of type %T", y, y)
    }
	return y
}
