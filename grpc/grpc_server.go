package up4net

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	sHost     = flag.String("host", "localhost", "Host name or ip address of server")
	sPort     = flag.Int("port", 10101, "The server port")
	sTLS      = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	sCertFile = flag.String("cert_file", "cert.pem", "The TLS cert file")
	sKeyFile  = flag.String("key_file", "cert.key", "The TLS key file")
)

// GRPCServerImpl interface definition of gRPC service implemention
type GRPCServerImpl interface {
	RegisterServer(*grpc.Server)
}

// RunGRPCServer run gRPC server
// ctx context to controll server stop
// impl your grpc server implemention
// customOpts custom options passing to gRPC server
func RunGRPCServer(ctx context.Context, impl GRPCServerImpl, customOpts ...grpc.ServerOption) error {
	// parse command line arguments
	if !flag.Parsed() {
		flag.Parse()
	}
	// fill server operations
	var opts []grpc.ServerOption
	if *sTLS {
		creds, err := credentials.NewServerTLSFromFile(*sCertFile, *sKeyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
			return fmt.Errorf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	opts = append(opts, customOpts...)
	// create gRPC server and regiter with sevice impl
	s := grpc.NewServer(opts...)
	impl.RegisterServer(s)
	// init listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *sHost, *sPort))
	if err != nil {
		log.Fatalf("Failed to listen at %s/%d, %v", *sHost, *sPort, err)
		return fmt.Errorf("Failed to listen at %s/%d, %v", *sHost, *sPort, err)
	}
	// serve in a go routine
	go s.Serve(listener)
	defer s.Stop()

	ctxWithCancel, cancel := context.WithCancel(ctx)
	// listen signal for graceful quit
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c)
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
				cancel()
			}
		}
	}()

	// block to listen quit signals
	select {
	case <-ctxWithCancel.Done():
		log.Println("Server exiting...")
	}
	return nil
}
