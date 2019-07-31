package server

import (
	"context"

	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *Server) JoinGame(ctx context.Context, req *proto.JoinGameRequest) (*proto.Player, error) {
	// Start by checking that the game we've requested is
	// actually a live game.
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.Player{},
			status.Errorf(codes.NotFound, "GameID %v not a valid game", req.GameID)
	}
	p, err := game.Join(req.Name)
	if err != nil {
		return nil, err
	}
	return p, nil
}
