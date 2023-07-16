package server

import (
	"context"
	pb "github.com/fuks-kit/doorman/proto"
	"github.com/fuks-kit/doorman/workspace"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type DoormanServer struct {
	pb.DoormanServer
}

func NewDoormanServer() *DoormanServer {
	return &DoormanServer{}
}

func (server *DoormanServer) CheckAccount(ctx context.Context, _ *emptypb.Empty) (*pb.AccountState, error) {
	token, err := verifyCredentials(ctx)
	if err != nil {
		return nil, err
	}

	user, err := authClient.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	access, err := workspace.HasOfficeAccess(token.UID, user.Email)
	if err != nil {
		return nil, err
	}

	// TODO: Add state infos

	return &pb.AccountState{
		HasAccess: access,
	}, nil
}

func (server *DoormanServer) OpenDoor(ctx context.Context, _ *emptypb.Empty) (*pb.DoorState, error) {
	log.Printf("OpenDoor")
	return &pb.DoorState{}, nil
}
