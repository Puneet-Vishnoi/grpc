package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	proto "github.com/Puneet-Vishnoi/grpc_tutorial_serverside_streaming/proto/hello"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientField struct {
	conn   *grpc.ClientConn
	client proto.ExampleClient
}


func ExampleClient() ClientField {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	client := proto.NewExampleClient(conn)
	return ClientField{
		conn:   conn,
		client: client,
	}
}

func (c *ClientField) Close() {
	c.conn.Close()
}

func main() {

	client := ExampleClient()

	r := gin.Default()
	r.GET("/sent", client.clientConnectionServer)

	r.Run(":8080")
}

func (cl *ClientField) clientConnectionServer(c *gin.Context) {
	req := &proto.HelloRequest{SomeString: "Request ker raha hu"}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := cl.client.ServerReply(ctx, req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalReply := 0
	for {
		message, err := stream.Recv()
		if err == io.EOF{
			break
		}

		fmt.Println("reply msg: "+ message.Reply)
		time.Sleep(1* time.Second)
		totalReply++
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "message sent succesfully to server",
		"totalReply": totalReply,
	})
}