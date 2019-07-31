package server

import (
	"context"

	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *Server) GameStatus(ctx context.Context, req *proto.GameStatusRequest) (*proto.GameStatusResponse, error) {
	// Start by checking that the game we've requested is
	// actually a live game.
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.GameStatusResponse{},
			status.Errorf(codes.NotFound, "GameID %v not a valid game", req.GameID)
	}

	return game.Status(req.PlayerID, req.PlayerSecret)
}
