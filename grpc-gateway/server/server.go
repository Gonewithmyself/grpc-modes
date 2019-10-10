package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync/atomic"

	"github.com/Gonewithmyself/grpc-modes/grpc-gateway/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

type echo struct {
}

var (
	endpoint = ":9090"
)

var x int64

func (e *echo) Echo(ctx context.Context, req *proto.Req) (*proto.Resp, error) {
	val := req.Value
	atomic.AddInt64(&x, 1)
	if "" == val {
		val = "null"
	}
	return &proto.Resp{Value: val + " count:" + strconv.Itoa(int(x))}, nil
}

func Serve() {
	grpcServer := grpc.NewServer()
	gwmux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: true, EmitDefaults: true}))
	proto.RegisterEchoServiceServer(grpcServer, new(echo))
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	if err := proto.RegisterEchoServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, opts); nil != err {
		panic(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	server := http.Server{
		Addr: endpoint,
		// Handler: util.GrpHandlerFunc(grpcServer, mux),
		Handler: mux,
	}

	l, err := net.Listen("tcp", endpoint)
	if nil != err {
		panic(err)
	}
	m := cmux.New(l)

	go server.Serve(m.Match(cmux.HTTP1Fast()))
	go grpcServer.Serve(m.Match(cmux.HTTP2()))

	fmt.Println("go")
	m.Serve()
}
