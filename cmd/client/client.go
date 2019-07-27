package main

import (
	"context"
	"io"
	"log"
	"time"

	desertproto "github.com/chronojam/solarium/pkg/gamemodes/desert/proto"
	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc"
)

func main() {
	grpc.EnableTracing = true
	conn, err := grpc.Dial(":8443", grpc.WithInsecure())
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
			for _, e := range not.Events {
				log.Printf("Global: %v", e.Notification)
			}
		}
	}()

	resp, err := client.NewGame(context.Background(), &proto.NewGameRequest{
		Gamemode:   "DesertPlanet",
		Difficulty: 1,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	time.Sleep(time.Second * 1)

	log.Printf("Description %v\nName: %v", resp.Description, resp.Name)

	chronojam, err := client.JoinGame(context.Background(), &proto.JoinGameRequest{
		GameID: resp.GameID,
		Name:   "Chronojam",
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	jemma, err := client.JoinGame(context.Background(), &proto.JoinGameRequest{
		GameID: resp.GameID,
		Name:   "Jemma",
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
			for _, e := range not.Events {
				log.Printf("%v", e)
				if e.DesertPlanet.DesertPlanetFailed != nil {
					log.Fatalf("You loose!, Final Score: %v", e.DesertPlanet.DesertPlanetFailed.Score)
				}
			}
		}
	}()

	go func() {
		for {
			// Occasionally get the status
			time.Sleep(time.Second * 3)
			_, err = client.GameStatus(context.Background(), &proto.GameStatusRequest{
				GameID: resp.GameID,
			})
			if err != nil {
				log.Fatalf("%v", err)
			}

			//log.Printf("%v", status.DesertPlanet)
		}
	}()

	// Gather some water
	go func() {
		for {
			time.Sleep(time.Second * 2)
			_, err = client.DoAction(context.Background(), &proto.DoActionRequest{
				PlayerID: chronojam.SecretKey,
				GameID:   resp.GameID,
				DesertPlanet: &desertproto.DesertPlanetAction{
					GatherWater: &desertproto.DesertPlanetGatherWater{},
				},
			})
			if err != nil {
				log.Fatalf("%v", err)
			}
		}
	}()

	for {
		time.Sleep(time.Second * 2)
		_, err = client.DoAction(context.Background(), &proto.DoActionRequest{
			PlayerID: jemma.SecretKey,
			GameID:   resp.GameID,
			DesertPlanet: &desertproto.DesertPlanetAction{
				GatherFood: &desertproto.DesertPlanetGatherFood{},
			},
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		time.Sleep(time.Second * 2)
		_, err = client.DoAction(context.Background(), &proto.DoActionRequest{
			PlayerID: jemma.SecretKey,
			GameID:   resp.GameID,
			DesertPlanet: &desertproto.DesertPlanetAction{
				GatherComponents: &desertproto.DesertPlanetGatherComponents{},
			},
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
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
