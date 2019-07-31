package server

import (
	"log"

	proto "github.com/chronojam/solarium/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
				return status.Errorf(codes.Canceled, "Stream Closed", err)
			}
		}
	}

	return nil
}
