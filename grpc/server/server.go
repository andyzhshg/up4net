package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Config config of gRPC seaver
type Config struct {
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	TLS         bool   `toml:"tls_on"`
	CertFile    string `toml:"cert_file"`
	KeyFile     string `toml:"key_file"`
	CatchSignal bool   `toml:"catch_signal"`
}

// DefaultConfig get a default config
func DefaultConfig() *Config {
	return &Config{
		Host:        "0.0.0.0",
		Port:        10101,
		TLS:         false,
		CertFile:    "cert.pem",
		KeyFile:     "cert.key",
		CatchSignal: false,
	}
}

// Register interface definition of gRPC service implemention
type Register interface {
	RegisterServer(*grpc.Server)
}

// Run run gRPC server
// ctx context to controll server stop
// impl your grpc server implemention
// cfg configuration of server
// customOpts custom options passing to gRPC server
func Run(ctx context.Context, reg Register, cfg *Config, customOpts ...grpc.ServerOption) error {
	// fill server operations
	var opts []grpc.ServerOption
	if cfg.TLS {
		creds, err := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
		if err != nil {
			return fmt.Errorf("Failed to generate credentials %v", err)
		}
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}
	opts = append(opts, customOpts...)
	// create gRPC server and regiter with sevice impl
	s := grpc.NewServer(opts...)
	reg.RegisterServer(s)
	// init listener
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		return fmt.Errorf("Failed to listen at %s/%d, %v", cfg.Host, cfg.Port, err)
	}
	// serve in a go routine
	go s.Serve(listener)
	defer s.Stop()

	ctxWithCancel, cancel := context.WithCancel(ctx)
	if cfg.CatchSignal {
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
	}

	// block to listen quit signals
	select {
	case <-ctxWithCancel.Done():
		log.Println("Server exiting...")
	}
	return nil
}
