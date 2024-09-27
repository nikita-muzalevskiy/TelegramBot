package main

import (
	pb "NovelService/protos"
	"context"
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	if err := godotenv.Load("variables.env"); err != nil {
		log.Print("No .env file found")
	}
}

var port = flag.Int("port", 50052, "The server port")

type server struct {
	pb.UnimplementedNovelServer
}

func (s *server) Channel1(ctx context.Context, in *pb.CallbackRequest) (*pb.CallbackReply, error) {
	fmt.Println(ctx)
	return MessageToCallback(in.GetUser(), in.GetParam(), messageList[in.GetAction()]), nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNovelServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
