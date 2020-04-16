# up4net

A golang micro service net utility library.


## gRPC

A scaffold to write a gRPC server or client in just a few lines of code.

We use `example/grpc` as an example.

### Generate `pb.go` file

```bash
cd example/grpc
protoc -I ./proto ./proto/hello.proto --go_out=plugins=grpc:proto
```

### gRPC Server

To run a gRPC server, you only have to implement an extra interface- `Register`, which has only one method to implement- `RegisterServer(*grpc.Server)`.

We call `RegisterHelloServer` in the generated `*.pb.go` file for this example. You should find your corresponding `RegisterXxxxx` in your own `*.pb.go` file.

```go
func (s *myHelloServer) RegisterServer(gSvr *grpc.Server) {
	proto.RegisterHelloServer(gSvr, s)
}
```

Now you can start a gRPC server with one line of code- 

```go
grpc_server.Run(context.TODO(), &myHelloServer{}, grpc_server.DefaultConfig())
```

#### Configure

If you don't want the defaut configuration provided by `grpc_server.DefaultConfig()`, you can pass your own config as argument.

#### Other ServerOption

You can configure all ServerOption(s) by passing parameter to `Run`, this will override default ServerOption.

### gRPC Client

To start a gRPC client, one has to configure and start a net connection. With the help of this library, you can do this in one line of code-

```go
grpc_client.NewClientConnection(grpc_client.DefaultConfig())
```

#### Configure

If you don't want the defaut configuration provided by `grpc_client.DefaultConfig()`, you can pass your own config as argument.


#### Other DialOption

You can configure all ClientConn(s) by passing parameter to `NewClientConnection`, this will override default ClientConn.
