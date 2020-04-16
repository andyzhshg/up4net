package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/andyzhshg/up4net/example/grpc/proto"
	grpc_client "github.com/andyzhshg/up4net/grpc/client"
)

func main() {
	conn, err := grpc_client.NewClientConnection(grpc_client.DefaultConfig())
	if err != nil {
		log.Fatalf("fail to get connection: %v", err)
	}
	defer conn.Close()
	cli := proto.NewHelloClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	for i := 0; i < 10; i++ {
		reply, _ := cli.SayHello(ctx, &proto.HelloRequest{Message: fmt.Sprintf("Hi, %d", i+1)})
		repStr, _ := json.Marshal(reply)
		log.Printf("Got reply: %s", repStr)
	}
}
