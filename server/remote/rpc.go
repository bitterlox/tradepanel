package remote

import (
	"context"
	pb "github.com/bitterlox/tradepanel/server/remote/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "[RPC] ", log.LstdFlags)

type Server struct {
	pb.UnimplementedTradingApiServer
}

func NewServer() *Server {
	return &Server{pb.UnimplementedTradingApiServer{}}
}

func (s *Server) Status(ctx context.Context, _ *pb.StatusRequest) (*pb.StatusResponse, error) {
	t := timestamppb.Now()

	log.Println("served status", t)

	return &pb.StatusResponse{
		Timestamp: t,
		Msg:       "hello from server fren, here is your time",
	}, nil
}
