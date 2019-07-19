package main

import (
	"context"
	"io"
	"log"

	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc"
)

func main() {
	grpc.EnableTracing = true
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	//defer conn.Close()

	client := proto.NewSolariumClient(conn)
	resp, err := client.JoinGame(context.Background(), &proto.JoinGameRequest{GameID: "12345"})
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("SecretKey: %v", resp.SecretKey)
	log.Printf("PlanetName: %v", resp.PlanetName)

	stream, err := client.GameUpdate(context.Background(), &proto.GameUpdateRequest{})
	if err != nil {
		log.Fatalf("%v", err)
	}

	for {
		not, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("%v", err)
		}

		log.Printf("%v", not.Notification)
	}
}
