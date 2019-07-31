package server

import (
	"context"

	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *Server) DoAction(ctx context.Context, req *proto.DoActionRequest) (*proto.DoActionResponse, error) {
	// Start by checking that the game we've requested is
	// actually a live game.
	game, ok := g.Games[req.GameID]
	if !ok {
		return &proto.DoActionResponse{},
			status.Errorf(codes.NotFound, "GameID %v not a valid game", req.GameID)
	}
	// Also check if we've got a playerId and a secret key, because if we dont then also
	// return an error.
	if req.GetPlayerID() == "" || req.GetPlayerSecret() == "" {
		return &proto.DoActionResponse{},
			status.Errorf(codes.PermissionDenied, "PlayerID and PlayerSecret are both required parameters")
	}

	// Parse the action, then pass that down to the gamemode itself.
	a, err := ParseAction(req)
	if err != nil {
		return &proto.DoActionResponse{}, err
	}
	err = game.PlayerDoAction(a, req.PlayerID, req.PlayerSecret)
	if err != nil {
		return nil, err
	}

	return &proto.DoActionResponse{}, nil
}
