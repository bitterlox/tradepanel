package backend

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "localhost:50051"
	defaultName = "wails-client"
)

type Backend struct {
	conn *grpc.ClientConn
}

func NewBackend() *Backend {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	return &Backend{conn: conn}
}

func (b *Backend) Greet(greeting string) string {
	c := pb.NewGreeterClient(b.conn)

	// Contact the server and print out its response.
	name := defaultName

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.SayHello(ctx, &pb.HelloRequest{Name: "default_"+name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: greeting})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	return r.GetMessage()
}
