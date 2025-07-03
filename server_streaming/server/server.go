package main

import (
	"fmt"
	"log"
	"net"

	proto "github.com/Puneet-Vishnoi/grpc_tutorial_serverside_streaming/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listner, tcpErr := net.Listen("tcp", ":9000") // Starts a TCP listener on port 9000. This means your server will listen for incoming network connections on port 9000.

	if tcpErr != nil {
		log.Fatalf("Failed to listen: %v", tcpErr)
	}

	srv := grpc.NewServer() //This creates a new gRPC server instance in Go. This is the actual server that will handle incoming gRPC requests.

	// Register service
	proto.RegisterExampleServer(srv, &server{})

	// Enable reflection for tools like grpcurl or Evans
	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		log.Fatalf("Failed to serve: %v", e)
	}
}

func (s *server) ServerReply(req *proto.HelloRequest, stream proto.Example_ServerReplyServer) error {
	fmt.Println("received reqest from client", req.SomeString)
	fmt.Println("hello from server")

	responses := []*proto.HelloResponse{
		{Reply: "Kya hua"},
		{Reply: "Sorry"},
		{Reply: "Man jao"},
	}

	for _, msg := range responses {
		err := stream.Send(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
