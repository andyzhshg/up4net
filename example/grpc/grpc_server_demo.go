package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/andyzhshg/up4net/example/grpc/grpc_demo"

	u4n_grpc "github.com/andyzhshg/up4net/grpc/server"
	"google.golang.org/grpc"
)

type myHelloServer struct {
}

func (s *myHelloServer) SayHello(ctx context.Context, req *grpc_demo.HelloRequest) (*grpc_demo.HelloReply, error) {
	reqStr, _ := json.Marshal(req)
	log.Printf("Got request: %s", reqStr)
	return &grpc_demo.HelloReply{Message: fmt.Sprintf("Got: %s", req.Message)}, nil
}

func (s *myHelloServer) RegisterServer(gSvr *grpc.Server) {
	grpc_demo.RegisterHelloServer(gSvr, s)
}

func main() {
	u4n_grpc.RunGRPCServer(context.TODO(), &myHelloServer{})
}
