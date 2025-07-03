package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	proto "github.com/Puneet-Vishnoi/grpc_tutorial_clientside_streaming/proto/hello"
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
	reqs := []*proto.HelloRequest{
		{SomeString: "Request 1"},
		{SomeString: "Request 2"},
		{SomeString: "Request 3"},
		{SomeString: "Request 4"},
		{SomeString: "Request 5"},
		{SomeString: "Request 6"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := cl.client.ServerReply(ctx)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs{
		err = stream.Send(req)
		if err != nil{
			fmt.Println("request not fulfill")
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// clone the stream and get the server response
	res, err := stream.CloseAndRecv()

	if err != nil{
		log.Printf("Error receiving response: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "message sent succesfully to server",
		"reply": res.Reply,
	})
}