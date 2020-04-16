package client

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Config config of gRPC client connect
type Config struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	TLS          bool   `toml:"tls_on"`
	CaFile       string `toml:"ca_file"`
	HostOverride string `toml:"host_override"`
}

// DefaultConfig get a default config
func DefaultConfig() *Config {
	return &Config{
		Host:         "127.0.0.1",
		Port:         10101,
		TLS:          false,
		CaFile:       "ca.pem",
		HostOverride: "www.example.com",
	}
}

// NewClientConnection create a new gRPC client
func NewClientConnection(cfg *Config, customOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if cfg.TLS {
		creds, err := credentials.NewClientTLSFromFile(cfg.CaFile, cfg.HostOverride)
		if err != nil {
			return nil, fmt.Errorf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, customOpts...)

	// Dial to server
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), opts...)
	if err != nil {
		return nil, fmt.Errorf("Fail to dial: %v", err)
	}
	return conn, nil
}
