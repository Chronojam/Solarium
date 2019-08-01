package server

import (
	proto "github.com/chronojam/solarium/proto"
)

type Gamemode interface {
	Setup()
	Description() string
	Join(name string) (*proto.Player, error)
	Simulate()
	PlayerDoAction(a interface{}, pid, secret string) error
	NextEvent() *proto.GameEvent
	Status(pid, psecret string) (*proto.GameStatusResponse, error)
	IsRunning() bool
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
		// Check if its the right game
		if id == idi {
			// send the event to all listeners.
			for _, lis := range listeners {
				lis <- event
			}
		}
	}
}
func (g *Server) DispatchToGlobal(e *proto.GlobalEvent) {
	for _, lis := range g.GlobalListeners {
		lis <- e
	}
}
