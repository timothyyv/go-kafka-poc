package grpc

import (
	"context"
	"fmt"

	pb "grule-demo/grpc/client"

	grpcgo "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Connector struct{}

func NewGRPCConnector() *Connector {
	return &Connector{}
}

func (c *Connector) Dial() (*grpcgo.ClientConn, error) {
	conn, err := grpcgo.Dial("localhost:9090", grpcgo.WithBlock(), grpcgo.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("could not dial grpc %w", err)
	}
	return conn, nil
}

func (c *Connector) SendMessage(ctx context.Context, message string) (*pb.AggregationResponse, error) {
	var res *pb.AggregationResponse
	conn, err := c.Dial()
	if err != nil {
		return nil, fmt.Errorf("could not dial: %w", err)
	}
	defer conn.Close()

	client := pb.NewAggregationClient(conn)
	res, err = client.SendAggregation(ctx, &pb.AggregationRequest{Input: message})
	if err != nil {
		responseStatus := status.Convert(err)
		if responseStatus == nil {

			return nil, fmt.Errorf("could not send message: %w", err)
		}

		return nil, fmt.Errorf("could not send message, received status %q: %s", responseStatus.Code(), responseStatus.Message())
	}

	return res, nil
}