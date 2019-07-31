package server

import (
	"github.com/chronojam/solarium/pkg/gamemodes/thewolfgame"
	proto "github.com/chronojam/solarium/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	AvaliableGameModes = map[string]func(diff proto.NewGameRequest_DifficultyLevel) Gamemode{
		proto.NewGameRequest_THEWOLFGAME.String(): func(diff proto.NewGameRequest_DifficultyLevel) Gamemode {
			return thewolfgame.New()
		},
	}
)

// Returns the inner action from a given request.
func ParseAction(req *proto.DoActionRequest) (interface{}, error) {
	var t interface{}
	switch {
	case req.TheWolfGame != nil:
		t = req.TheWolfGame
	default:
		return nil, status.Errorf(codes.InvalidArgument, "Requested action is not avaliable")
	}

	return t, nil
}
