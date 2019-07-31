package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"time"

	wolfproto "github.com/chronojam/solarium/pkg/gamemodes/thewolfgame/proto"
	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc"
)

var (
	pMap = []*proto.Player{}
)

func joinNewPlayer(gid, name string, client proto.SolariumClient) {
	p, err := client.JoinGame(context.Background(), &proto.JoinGameRequest{
		GameID: gid,
		Name:   name,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}

	pMap = append(pMap, p)
}

func main() {
	grpc.EnableTracing = true
	conn, err := grpc.Dial(":8443", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("%v", err)
	}
	//defer conn.Close()

	client := proto.NewSolariumClient(conn)
	resp, err := client.NewGame(context.Background(), &proto.NewGameRequest{
		Gamemode:   proto.NewGameRequest_THEWOLFGAME,
		Difficulty: 1,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	joinNewPlayer(resp.ID, "Chronojam", client)
	joinNewPlayer(resp.ID, "Chronojam1", client)
	joinNewPlayer(resp.ID, "Chronojam2", client)
	joinNewPlayer(resp.ID, "Chronojam3", client)
	joinNewPlayer(resp.ID, "Chronojam4", client)
	joinNewPlayer(resp.ID, "Chronojam5", client)
	joinNewPlayer(resp.ID, "Chronojam6", client)
	joinNewPlayer(resp.ID, "Chronojam7", client)
	joinNewPlayer(resp.ID, "Chronojam8", client)
	joinNewPlayer(resp.ID, "Chronojam9", client)

	for _, p := range pMap {
		log.Printf("Registered Player: %v", p.Name)
	}
	// Get Generic game events.
	go func() {
		stream, err := client.GameUpdate(context.Background(), &proto.GameUpdateRequest{
			GameID: resp.ID,
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
			}
		}
	}()

	// Get everyone to vote to start
	for _, p := range pMap {
		client.DoAction(context.Background(), &proto.DoActionRequest{
			PlayerID:     p.ID,
			PlayerSecret: p.Secret,
			GameID:       resp.ID,
			TheWolfGame: &wolfproto.TheWolfGameAction{
				StartVote: &wolfproto.TheWolfGameAction_VoteStart{},
			},
		})
		log.Printf("Player %v casted start vote", p.Name)
	}

	wolves := []*wolfproto.TheWolfGamePlayer{}
	villagers := []*wolfproto.TheWolfGamePlayer{}
	for _, p := range pMap {
		// Now everyone checks if they are a werewolf.
		me, _ := client.GameStatus(context.Background(), &proto.GameStatusRequest{
			GameID:       resp.ID,
			PlayerID:     p.ID,
			PlayerSecret: p.Secret,
		})

		if me.TheWolfGame.Players[0].Role == wolfproto.TheWolfGamePlayer_WEREWOLF {
			wolves = append(wolves, me.TheWolfGame.Players[0])
			log.Printf("%v is a werewolf", p.Name)
		} else {
			villagers = append(villagers, me.TheWolfGame.Players[0])
		}
	}

	vindex := 0
	for {
		// Get the state of the game
		status, err := client.GameStatus(context.Background(), &proto.GameStatusRequest{
			GameID: resp.ID,
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		if !status.TheWolfGame.IsStarted {
			continue
		}
		log.Printf("Preparing the lynch")

		// Choose someone to get killed
		ded := villagers[vindex]
		vindex++

		// Everyone votes the same for testings.
		for _, p := range pMap {
			_, _ = client.DoAction(context.Background(), &proto.DoActionRequest{
				PlayerID:     p.ID,
				PlayerSecret: p.Secret,
				GameID:       resp.ID,
				TheWolfGame: &wolfproto.TheWolfGameAction{
					Vote: &wolfproto.TheWolfGameAction_VoteMurder{
						PlayerId: ded.ID,
					},
				},
			})

			log.Printf("%v has voted", p.ID)
		}

		time.Sleep(10 * time.Second)
		log.Printf("Fetching new status")
		stat, _ := client.GameStatus(context.Background(), &proto.GameStatusRequest{
			GameID: resp.ID,
		})
		b, _ := json.MarshalIndent(stat, "", "  ")
		log.Printf("A new round is starting!")
		log.Printf("%v", string(b))
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
