package server

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chronojam/solarium/pkg/game/interfaces"
	"github.com/chronojam/solarium/pkg/gamemodes/desert-planet"

	"github.com/google/uuid"

	proto "github.com/chronojam/solarium/proto"
)

var (
	AvaliableGameModes = map[string]func(diff int) Gamemode{
		"DesertPlanet": func(diff int) Gamemode {
			return desert.New(diff)
		},
	}
)

type Gamemode interface {
	Setup()
	NextEvent() interfaces.GameEvent
	Simulate()
	Join(name string) (interfaces.Player, error)
	PlayerDoAction(pid, action string) error
}

type Server struct {
	// For now we'll store the game info here
	// and assume we can only host a single game.
	Games map[string]Gamemode

	History []string

	// GameID: []chan string
	Listeners map[string]map[string]chan string
}

func New() *Server {
	s := &Server{
		Games: map[string]Gamemode{},
		Listeners: map[string]map[string]chan string{
			"Global": map[string]chan string{},
		},
	}
	//go s.NotificationSpreader()
	return s
}

/*func (g *Server) NotificationSpreader() {
	// Listen for notifications
	for {
		select {
		case notification := <-g.System.NotificationChan:
			log.Printf("Got Notification: %v", notification)
			// Put one copy onto the history list
			g.History = append(g.History, notification)

			// Send one to each listening channel.
			for _, c := range g.Listeners {
				c <- notification
			}
		}
	}
}*/

func (g *Server) DispatchToGameID(idi, message string) {
	for id, listeners := range g.Listeners {
		if id == idi {
			for _, lis := range listeners {
				go func() {
					lis <- message
				}()
			}
		}
	}
}

func (g *Server) NewGame(ctx context.Context, req *proto.NewGameRequest) (*proto.NewGameResponse, error) {
	gm, ok := AvaliableGameModes[req.Gamemode]
	if !ok {
		return &proto.NewGameResponse{}, nil
	}
	guid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	g.Listeners[guid.String()] = map[string]chan string{}
	g.Games[guid.String()] = gm(int(req.Difficulty))
	g.DispatchToGameID("Global", fmt.Sprintf("A new game has started of: %v, with id: %v", req.Gamemode, guid.String()))
	g.StartGame(guid.String())
	return &proto.NewGameResponse{
		GameID: guid.String(),
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
		SecretKey: p.Id(),
	}, nil
}

func (g *Server) StartGame(id string) {
	game, ok := g.Games[id]
	if !ok {
		return
	}
	go func() {
		log.Printf("Doing Setup()")
		game.Setup()
		for {
			time.Sleep(time.Second * 5)
			game.Simulate()
		}
	}()
	go func() {
		for {
			log.Printf("Waiting for event..")
			e := game.NextEvent()
			log.Printf(e.Description)
			g.DispatchToGameID(id, e.Description)
		}
	}()
}

func (g *Server) DoAction(ctx context.Context, req *proto.DoActionRequest) (*proto.DoActionResponse, error) {
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.DoActionResponse{}, nil
	}
	err := game.PlayerDoAction(req.SecretKey, req.Action)
	if err != nil {
		return nil, err
	}
	return &proto.DoActionResponse{}, nil
}

func (g *Server) GlobalUpdate(req *proto.GlobalUpdateRequest, stream proto.Solarium_GlobalUpdateServer) error {
	me := make(chan string)
	cid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	cuuid := cid.String()
	g.Listeners["Global"][cuuid] = me
	for {
		select {
		case notification := <-me:
			if err := stream.Send(&proto.GlobalUpdateResponse{Notification: notification}); err != nil {
				log.Printf("%v", err)
				delete(g.Listeners["Global"], cuuid)
				return nil
			}
		}
	}

	return nil
}

func (g *Server) GameUpdate(req *proto.GameUpdateRequest, stream proto.Solarium_GameUpdateServer) error {
	// Keep this guy open. Continually read from the notification channel and send it off
	// to whoever is connected.
	me := make(chan string)
	cid, err := uuid.NewRandom()
	if err != nil {
		return err
	}
	cuuid := cid.String()
	g.Listeners[req.GameID][cuuid] = me
	log.Printf("Stream %v", cuuid)

	// Continully send updates to the client as long as this connection is open.
	for {
		select {
		case notification := <-me:
			if err := stream.Send(&proto.GameUpdateResponse{Notification: notification}); err != nil {
				log.Printf("%v", err)
				delete(g.Listeners[req.GameID], cuuid)
				return nil
			}
		}
	}

	return nil
}
