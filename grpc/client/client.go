package main

import (
	"context"
	"log"
	"net/http"
	"time"

	proto "github.com/Puneet-Vishnoi/grpc_tutorial/proto/hello"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientField struct {
	conn   *grpc.ClientConn
	client proto.EmapleClient
}

func EmapleClient() ClientField {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	client := proto.NewEmapleClient(conn)
	return ClientField{
		conn:   conn,
		client: client,
	}
}

func (c *ClientField) Close() {
	c.conn.Close()
}

func main() {

	client := EmapleClient()

	r := gin.Default()
	r.GET("/sent-message-to-Server/:message", client.clientConnectionServer)

	r.Run(":8080")
}

func (cl *ClientField) clientConnectionServer(c *gin.Context) {
	message := c.Param("message")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &proto.HelloRequest{SomeString: message}

	res, err := cl.client.ServerReply(ctx, req)
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "message sent succesfully to server " + message,
		"reply": res.Reply,
	})
}
