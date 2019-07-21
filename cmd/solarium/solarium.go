package main

import (
	"log"
	"net"

	"github.com/chronojam/solarium/pkg/server"
	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc"
)

func main() {
	grpc.EnableTracing = true
	conn, err := net.Listen("tcp", "0.0.0.0:8443")
	if err != nil {
		log.Fatalf(err.Error())
	}

	g := grpc.NewServer()
	proto.RegisterSolariumServer(g, server.New())

	log.Fatalf("%v", g.Serve(conn))
}
