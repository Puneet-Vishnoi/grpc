package main

import (
	"context"
	"fmt"
	"log"
	"net"

	proto "github.com/Puneet-Vishnoi/grpc_tutorial/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedEmapleServer
}

func main() {
	listner, tcpErr := net.Listen("tcp", ":9000")// Starts a TCP listener on port 90000. This means your server will listen for incoming network connections on port 9000.

	if tcpErr != nil {
		log.Fatalf("Failed to listen: %v", tcpErr)
	}

	srv := grpc.NewServer() //This creates a new gRPC server instance in Go. This is the actual server that will handle incoming gRPC requests.
	
	// Register service
	proto.RegisterEmapleServer(srv, &server{})

	// Enable reflection for tools like grpcurl or Evans
	reflection.Register(srv)

	if e := srv.Serve(listner); e!=nil{
		log.Fatalf("Failed to serve: %v", e)
	}
}


func (s *server) ServerReply(c context.Context, req *proto.HelloRequest)(*proto.HelloResponse, error){
	fmt.Println("received reqest from client", req.SomeString)
	fmt.Println("hello from server")

	return &proto.HelloResponse{Reply: "hello ji"}, nil
}