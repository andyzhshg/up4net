# up4net

A golang micro service net utility library.


## gRPC

A scaffold to write a gRPC server or client in just a few lines of code.

We use `example/grpc` as an example.

### Generate `pb.go` file

```bash
cd example/grpc
mkdir grpc_demo
protoc -I ./ ./grpc_demo.proto --go_out=plugins=grpc:grpc_demo
```

### gRPC Server

To run a gRPC server, you only have to implement an extra interface- `GRPCServerImpl`, which has only one method to implement- `RegisterServer(*grpc.Server)`.

We call `RegisterHelloServer` in the generated `*.pb.go` file for this example. You should find your corresponding `RegisterXxxxx` in your own `*.pb.go` file.

```go
func (s *myHelloServer) RegisterServer(gSvr *grpc.Server) {
	grpc_demo.RegisterHelloServer(gSvr, s)
}
```

Now you can start a gRPC server with one line of code- 

```go
u4n_grpc.RunGRPCServer(context.TODO(), &myHelloServer{})
```

#### Command line arguments

We predefined some command line arguments, you can pass these arguments to override the default ones.

|name|default|description|
|----|-------|-----------|
|--host|localhost|Host name or ip address of server|
|--port|10101|The server port|
|--tls|false|Connection uses TLS if true, else plain TCP|
|--cert_file|cert.pem|The TLS cert file|
|--key_file|cert.key|The TLS key file|

#### Other ServerOption

You can configure all ServerOption(s) by passing parameter to `RunGRPCServer`, this will override default ServerOption.

### gRPC Client

To start a gRPC client, one has to configure and start a net connection. With the help of this library, you can do this in one line of code-

```go
conn, err := u4n_grpc.NewGRPCClientConnection()
```

#### Command line arguments

We predefined some command line arguments, you can pass these arguments to override the default ones.

|name|default|description|
|----|-------|-----------|
|--host_addr|localhost:10101|The server address in the format of host:port|
|--tls_on|false|Connection uses TLS if true, else plain TCP|
|--ca_file|cert.pem|The file containing the CA root cert file|
|--host_override|www.example.com|The server name use to verify the hostname returned by TLS handshake|

#### Other DialOption

You can configure all ServerOption(s) by passing parameter to `RunGRPCServer`, this will override default ServerOption.