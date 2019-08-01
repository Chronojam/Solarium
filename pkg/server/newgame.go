package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chronojam/solarium/pkg/namegenerator"
	proto "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *Server) NewGame(ctx context.Context, req *proto.NewGameRequest) (*proto.SolariumGame, error) {
	gm, ok := AvaliableGameModes[req.Gamemode.String()]
	if !ok {
		// Requested an invalid gamemode.
		return &proto.SolariumGame{}, status.Errorf(codes.InvalidArgument, "Gamemode: %v is not a valid gamemode", req.Gamemode.String())
	}
	// GameID
	guid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	log.Printf("%v gid", guid.String())
	g.Listeners[guid.String()] = map[string]chan *proto.GameEvent{}
	newGame := gm(req.Difficulty)
	g.Games[guid.String()] = newGame
	g.DispatchToGlobal(&proto.GlobalEvent{
		Notification: fmt.Sprintf("A new game has started of: %v, with id: %v", req.Gamemode, guid.String()),
	})
	g.StartGame(newGame, guid.String())
	return &proto.SolariumGame{
		ID:          guid.String(),
		Description: newGame.Description(),
		Name:        namegenerator.GenerateNew(),
	}, nil
}

func (g *Server) StartGame(game Gamemode, id string) {
	go func() {
		game.Setup()
		game.Simulate()
		log.Printf("Cleaning Up")
		// Cleanup after game is done.
		// but wait 10seconds so people can get the latest updates
		time.Sleep(10 * time.Second)
		delete(g.Games, id)
		delete(g.Listeners, id)
	}()
	go func() {
		for {
			e := game.NextEvent()
			log.Printf("Game has emitted an event, sending dispatching to listeners...")
			g.DispatchToGameID(id, e)
		}
	}()
}
