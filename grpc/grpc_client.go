package up4net

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var (
	cAddr         = flag.String("host_addr", "localhost:10101", "The server address in the format of host:port")
	cTLS          = flag.Bool("tls_on", false, "Connection uses TLS if true, else plain TCP")
	cCaFile       = flag.String("ca_file", "ca.pem", "The file containing the CA root cert file")
	cHostOverride = flag.String("host_override", "www.example.com", "The server name use to verify the hostname returned by TLS handshake")
)

// NewGRPCClientConnection create a new gRPC client
func NewGRPCClientConnection(customOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	// parse command line arguments
	if !flag.Parsed() {
		flag.Parse()
	}
	var opts []grpc.DialOption
	if *cTLS {
		creds, err := credentials.NewClientTLSFromFile(*cCaFile, *cHostOverride)
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
			return nil, fmt.Errorf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	opts = append(opts, customOpts...)

	// Dial to server
	conn, err := grpc.Dial(*cAddr, opts...)
	if err != nil {
		log.Fatalf("Fail to dial: %v", err)
		return nil, fmt.Errorf("Fail to dial: %v", err)
	}
	return conn, nil
}
