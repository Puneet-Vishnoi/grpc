package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"

	proto "github.com/Puneet-Vishnoi/grpc_tutorial_clientside_streaming/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listner, tcpErr := net.Listen("tcp", ":9000")// Starts a TCP listener on port 90000. This means your server will listen for incoming network connections on port 9000.

	if tcpErr != nil {
		log.Fatalf("Failed to listen: %v", tcpErr)
	}

	srv := grpc.NewServer() //This creates a new gRPC server instance in Go. This is the actual server that will handle incoming gRPC requests.
	
	// Register service
	proto.RegisterExampleServer(srv, &server{})
	
	// Enable reflection for tools like grpcurl or Evans
	reflection.Register(srv)

	if e := srv.Serve(listner); e != nil {
		log.Fatal("Failed to serve: %v", e)
	}
}


func (s *server) ServerReply(stream proto.Example_ServerReplyServer)error{
	totalMessages := 0

	for{
		request, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa((totalMessages)),
			})
		}
		if err != nil{
			return nil
		}

		totalMessages++

		fmt.Println(request.SomeString)
	}
}