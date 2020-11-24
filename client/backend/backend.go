package backend

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

func (b *Backend) Greet(greeting string) string {
	c := pb(b.conn)

	// Contact the server and print out its response.
	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.SayHello(ctx, &pb.HelloRequest{Name: "default_" + name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: greeting})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.GetMessage()
}
