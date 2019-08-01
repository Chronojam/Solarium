package main

import (
	"bufio"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

	wolfproto "github.com/chronojam/solarium/pkg/gamemodes/thewolfgame/proto"
	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc"
)

var (
	pMap   = []*proto.Player{}
	GameID = "26598f3d-c358-4a13-8476-afde20b08a2e"
)

func joinNewPlayer(name string, client proto.SolariumClient) {
	p, err := client.JoinGame(context.Background(), &proto.JoinGameRequest{
		GameID: GameID,
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
	/*resp, err := client.NewGame(context.Background(), &proto.NewGameRequest{
		Gamemode:   proto.NewGameRequest_THEWOLFGAME,
		Difficulty: 1,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}*/
	//GameID = resp.ID
	GameID = "c060ae56-d8f5-41e8-8d4c-fed538a9f983"
	joinNewPlayer("James", client)
	joinNewPlayer("Jenna", client)
	joinNewPlayer("Kylie", client)
	joinNewPlayer("Fergus", client)
	joinNewPlayer("BoggyPete", client)
	joinNewPlayer("Nicola", client)
	joinNewPlayer("xXPvpGodXx", client)
	joinNewPlayer("Toestomper", client)
	joinNewPlayer("Applepresser", client)
	joinNewPlayer("Delimeats", client)

	for _, p := range pMap {
		log.Printf("Registered Player: %v", p.Name)
	}
	// Get Generic game events.
	go func() {
		stream, err := client.GameUpdate(context.Background(), &proto.GameUpdateRequest{
			GameID: GameID,
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
				if e.IsGameOver {
					// end the game.
					log.Fatalf("Game Over!")
				}
			}
		}
	}()

	// Get everyone to vote to start
	log.Printf("Waiting for input to start voteStart()")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	for _, p := range pMap {
		client.DoAction(context.Background(), &proto.DoActionRequest{
			PlayerID:     p.ID,
			PlayerSecret: p.Secret,
			GameID:       GameID,
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
			GameID:       GameID,
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

	for {
		// Get the state of the game
		status, err := client.GameStatus(context.Background(), &proto.GameStatusRequest{
			GameID: GameID,
		})
		if err != nil {
			log.Fatalf("%v", err)
		}
		if !status.TheWolfGame.IsStarted {
			continue
		}
		log.Printf("Preparing the lynch")

		// Choose someone to get killed
		vindex := 0
		for i, p := range villagers {
			if p.IsAlive {
				vindex = i
				p.IsAlive = !p.IsAlive
				break
			}
		}
		v := villagers[vindex]
		log.Printf("Waiting for input to vote.")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()

		// Everyone votes the same for testings.
		for _, p := range pMap {
			_, _ = client.DoAction(context.Background(), &proto.DoActionRequest{
				PlayerID:     p.ID,
				PlayerSecret: p.Secret,
				GameID:       GameID,
				TheWolfGame: &wolfproto.TheWolfGameAction{
					Vote: &wolfproto.TheWolfGameAction_VoteMurder{
						PlayerId: v.ID,
					},
				},
			})

			log.Printf("%v has voted", p.ID)
		}

		log.Printf("Fetching new status")
		stat, _ := client.GameStatus(context.Background(), &proto.GameStatusRequest{
			GameID: GameID,
		})
		_, _ = json.MarshalIndent(stat, "", "  ")
		log.Printf("A new round is starting!")
		//log.Printf("%v", string(b))
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
