package server

import (
	"context"
	"log"

	"github.com/google/uuid"

	"github.com/chronojam/solarium/pkg/planet"

	"github.com/chronojam/solarium/pkg/system"
	proto "github.com/chronojam/solarium/proto"
)

type Server struct {
	// For now we'll store the game info here
	// and assume we can only host a single game.
	System *system.System

	History   []string
	Listeners []chan string
}

func New() *Server {
	system := system.New()

	s := &Server{
		System: system,
	}
	go s.NotificationSpreader()
	go system.Simulate()
	return s
}

func (g *Server) NotificationSpreader() {
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
}

func (g *Server) JoinGame(ctx context.Context, req *proto.JoinGameRequest) (*proto.JoinGameResponse, error) {
	// Create a new body for this player, and give them a special token so they can interact with it.
	log.Printf("Hello")
	secretKey, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	p := planet.NewRandom("", secretKey.String())
	g.System.AddBody(p)

	return &proto.JoinGameResponse{
		SecretKey:  secretKey.String(),
		PlanetName: p.Name(),
	}, nil
}

func (g *Server) GameUpdate(req *proto.GameUpdateRequest, stream proto.Solarium_GameUpdateServer) error {
	// Keep this guy open. Continually read from the notification channel and send it off
	// to whoever is connected.
	me := make(chan string)
	g.Listeners = append(g.Listeners, me)

	// Send all the current history first.
	for _, h := range g.History {
		go func() {
			me <- h
		}()
	}

	// Continully send updates to the client as long as this connection is open.
	for {
		select {
		case notification := <-me:
			if err := stream.Send(&proto.GameUpdateResponse{Notification: notification}); err != nil {
				log.Printf("%v", err)
			}
		}
	}

	return nil
}
