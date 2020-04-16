package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/andyzhshg/up4net/example/grpc/proto"

	grpc_server "github.com/andyzhshg/up4net/grpc/server"
	"google.golang.org/grpc"
)

type myHelloServer struct {
}

func (s *myHelloServer) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	reqStr, _ := json.Marshal(req)
	log.Printf("Got request: %s", reqStr)
	return &proto.HelloReply{Message: fmt.Sprintf("Got: %s", req.Message)}, nil
}

func (s *myHelloServer) RegisterServer(gSvr *grpc.Server) {
	proto.RegisterHelloServer(gSvr, s)
}

func main() {
	grpc_server.Run(context.TODO(), &myHelloServer{}, grpc_server.DefaultConfig())
}
