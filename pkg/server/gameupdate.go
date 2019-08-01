package server

import (
	"log"

	proto "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *Server) GameUpdate(req *proto.GameUpdateRequest, stream proto.Solarium_GameUpdateServer) error {
	// Keep this guy open. Continually read from the notification channel and send it off
	// to whoever is connected.
	if _, ok := g.Games[req.GameID]; !ok {
		// Game doesnt exist, or never existed.
		if err := stream.Send(&proto.GameUpdateResponse{Events: []*proto.GameEvent{
			&proto.GameEvent{
				Desc:       "This game doesnt exist, or is over!",
				IsGameOver: true,
			},
		}}); err != nil {
			return err
		}
		return nil
	}
	me := make(chan *proto.GameEvent, 10)
	cid, err := uuid.NewRandom()
	if err != nil {
		return status.Errorf(codes.Internal, "Oh No")
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
				return status.Errorf(codes.Canceled, "Stream Closed", err)
			}
		}
	}

	return nil
}
