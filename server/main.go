package main

import (
	"github.com/bitterlox/tradepanel/server/remote"
	pb "github.com/bitterlox/tradepanel/server/remote/proto"
	"github.com/pelletier/go-toml"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

var logger = log.New(os.Stdout, "[MAIN] ", log.LstdFlags)

type Config struct {
	Rpc remote.Config
}

func parseConfig(path string) (*Config, error) {

	var cfg Config

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = toml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func main() {

	cfg, err := parseConfig("config.toml")
	if err != nil {
		log.Fatal("error reading config: ", err)
	}

	logger.Printf("Printing config: %+v", cfg)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTradingApiServer(s, remote.NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	os.Exit(1)
}
