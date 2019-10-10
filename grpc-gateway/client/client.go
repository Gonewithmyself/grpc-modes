package client

import (
	"log"

	"github.com/Gonewithmyself/grpc-modes/grpc-gateway/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var endpoint = ":9090"

func Send(msg string) {
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Println(err)
	}

	c := proto.NewEchoServiceClient(conn)
	context := context.Background()
	body := &proto.Req{
		Value: msg,
	}

	r, err := c.Echo(context, body)
	if err != nil {
		log.Println(err)
	}

	log.Println(r.Value)
}
