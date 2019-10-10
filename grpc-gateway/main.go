package main

import (
	"os"

	"github.com/Gonewithmyself/grpc-modes/grpc-gateway/client"
	"github.com/Gonewithmyself/grpc-modes/grpc-gateway/server"
)

func main() {
	if len(os.Args) > 1 {
		client.Send(os.Args[1])
	} else {
		server.Serve()
	}
}
