package server

import (
	"context"
	pb "github.com/fuks-kit/doorman/proto"
	fuks "github.com/fuks-kit/doorman/workspace"
	"google.golang.org/protobuf/types/known/emptypb"
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

	access, err := fuks.HasOfficeAccess(token.UID, user.Email)
	if err != nil {
		return nil, err
	}

	return &pb.AccountState{
		HasAccess: access,
	}, nil
}
