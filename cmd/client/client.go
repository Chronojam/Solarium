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

	go func() {
		stream, err := client.GlobalUpdate(context.Background(), &proto.GlobalUpdateRequest{})
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
			log.Printf("Global: %v", not.Notification)
		}
	}()

	resp, err := client.NewGame(context.Background(), &proto.NewGameRequest{
		Gamemode:   "Desert-Planet",
		Difficulty: 1,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	jg, err := client.JoinGame(context.Background(), &proto.JoinGameRequest{
		GameID: resp.GameID,
		Name:   "Chronojam",
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	go func() {
		stream, err := client.GameUpdate(context.Background(), &proto.GameUpdateRequest{
			GameID: resp.GameID,
		})
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
			log.Printf("GID:%v --- %v", resp.GameID, not.Notification)
		}
	}()

	_, err = client.DoAction(context.Background(), &proto.DoActionRequest{
		SecretKey: jg.SecretKey,
		GameID:    resp.GameID,
		Action:    "me:GetWater",
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	for {
	}

	/*
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
	*/
}
