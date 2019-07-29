package server

import (
	"context"
	"fmt"
	"log"

	desert "github.com/chronojam/solarium/pkg/gamemodes/desert-planet"
	"github.com/chronojam/solarium/pkg/gamemodes/thewolfgame"
	"github.com/chronojam/solarium/pkg/namegenerator"

	"github.com/google/uuid"

	proto "github.com/chronojam/solarium/proto"
)

var (
	AvaliableGameModes = map[string]func(diff proto.NewGameRequest_DifficultyLevel) Gamemode{
		proto.NewGameRequest_DESERTPLANET.String(): func(diff proto.NewGameRequest_DifficultyLevel) Gamemode {
			return desert.New(diff)
		},
		proto.NewGameRequest_THEWOLFGAME.String(): func(diff proto.NewGameRequest_DifficultyLevel) Gamemode {
			return thewolfgame.New()
		},
	}
)

type Gamemode interface {
	Setup()
	Description() string
	Join(name string) (*proto.Player, error)
	Simulate()
	PlayerDoAction(req *proto.DoActionRequest) error
	NextEvent() *proto.GameEvent
	Status() *proto.GameStatusResponse
}

type Server struct {
	// For now we'll store the game info here
	// and assume we can only host a single game.
	Games map[string]Gamemode

	History []string

	// GameID: []chan string
	Listeners       map[string]map[string]chan *proto.GameEvent
	GlobalListeners map[string]chan *proto.GlobalEvent
}

func New() *Server {
	s := &Server{
		Games:           map[string]Gamemode{},
		Listeners:       map[string]map[string]chan *proto.GameEvent{},
		GlobalListeners: map[string]chan *proto.GlobalEvent{},
	}
	return s
}

func (g *Server) DispatchToGameID(idi string, event *proto.GameEvent) {
	for id, listeners := range g.Listeners {
		if id == idi {
			for _, lis := range listeners {
				go func() {
					lis <- event
				}()
			}
		}
	}
}

func (g *Server) DispatchToGlobal(e *proto.GlobalEvent) {
	for _, lis := range g.GlobalListeners {
		go func() {
			lis <- e
		}()
	}
}

func (g *Server) NewGame(ctx context.Context, req *proto.NewGameRequest) (*proto.NewGameResponse, error) {
	gm, ok := AvaliableGameModes[req.Gamemode.String()]
	if !ok {
		return &proto.NewGameResponse{}, nil
	}
	guid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	g.Listeners[guid.String()] = map[string]chan *proto.GameEvent{}
	newGame := gm(req.Difficulty)
	g.Games[guid.String()] = newGame
	g.DispatchToGlobal(&proto.GlobalEvent{
		Notification: fmt.Sprintf("A new game has started of: %v, with id: %v", req.Gamemode, guid.String()),
	})
	g.StartGame(guid.String())
	return &proto.NewGameResponse{
		GameID:      guid.String(),
		Description: newGame.Description(),
		Name:        namegenerator.GenerateNew(),
	}, nil
}

func (g *Server) JoinGame(ctx context.Context, req *proto.JoinGameRequest) (*proto.JoinGameResponse, error) {
	// Create a new body for this player, and give them a special token so they can interact with it.
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.JoinGameResponse{}, nil
	}
	p, err := game.Join(req.Name)
	if err != nil {
		return nil, err
	}
	return &proto.JoinGameResponse{
		SecretKey: p.ID,
	}, nil
}

func (g *Server) StartGame(id string) {
	game, ok := g.Games[id]
	if !ok {
		return
	}
	go func() {
		game.Setup()
		game.Simulate()
		log.Printf("Cleaning Up")
		// Cleanup after game is done.
		delete(g.Games, id)
		delete(g.Listeners, id)
	}()
	go func() {
		for {
			e := game.NextEvent()
			g.DispatchToGameID(id, e)
		}
	}()
}

func (g *Server) DoAction(ctx context.Context, req *proto.DoActionRequest) (*proto.DoActionResponse, error) {
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.DoActionResponse{}, nil
	}
	// Try and figure out which option we have sent
	// we should only ever send one, so just take the first we find.
	err := game.PlayerDoAction(req)
	if err != nil {
		return nil, err
	}

	return &proto.DoActionResponse{}, nil
}

func (g *Server) GameStatus(ctx context.Context, req *proto.GameStatusRequest) (*proto.GameStatusResponse, error) {
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.GameStatusResponse{}, nil
	}

	return game.Status(), nil
}

func (g *Server) GlobalUpdate(req *proto.GlobalUpdateRequest, stream proto.Solarium_GlobalUpdateServer) error {
	me := make(chan *proto.GlobalEvent)
	cid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	cuuid := cid.String()
	g.GlobalListeners[cuuid] = me
	for {
		select {
		case notification := <-me:
			if err := stream.Send(&proto.GlobalUpdateResponse{Events: []*proto.GlobalEvent{
				notification,
			}}); err != nil {
				log.Printf("%v", err)
				delete(g.GlobalListeners, cuuid)
				return nil
			}
		}
	}

	return nil
}

func (g *Server) GameUpdate(req *proto.GameUpdateRequest, stream proto.Solarium_GameUpdateServer) error {
	// Keep this guy open. Continually read from the notification channel and send it off
	// to whoever is connected.
	if _, ok := g.Games[req.GameID]; !ok {
		// Game doesnt exist, or never existed.
		if err := stream.Send(&proto.GameUpdateResponse{Events: []*proto.GameEvent{
			&proto.GameEvent{
				Desc:       "This game is already over!",
				IsGameOver: true,
			},
		}}); err != nil {
			return err
		}
		return nil
	}
	me := make(chan *proto.GameEvent)
	cid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	cuuid := cid.String()
	g.Listeners[req.GameID][cuuid] = me

	// Continully send updates to the client as long as this connection is open.
	// For now, dont bulk send
	for {
		select {
		case event := <-me:
			if err := stream.Send(&proto.GameUpdateResponse{Events: []*proto.GameEvent{event}}); err != nil {
				log.Printf("%v", err)
				delete(g.Listeners[req.GameID], cuuid)
				return nil
			}
		}
	}

	return nil
}
