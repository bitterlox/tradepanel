package client

import (
	"context"
	"log"
	"time"

	pb "github.com/bitterlox/tradepanel/server/remote/proto"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "wails-client"
)

type Backend struct {
	conn *grpc.ClientConn
}

func NewBackend() (*Backend, error) {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	return &Backend{conn: conn}, nil
}

func (b *Backend) Greet() string {
	c := pb.NewTradingApiClient(b.conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := c.Status(ctx, &pb.StatusRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return res.String()
}
